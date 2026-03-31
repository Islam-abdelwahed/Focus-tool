package blocker

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	hostsPath  = `C:\Windows\System32\drivers\etc\hosts`
	blockStart = "# === FOCUS BLOCK START ==="
	blockEnd   = "# === FOCUS BLOCK END ==="
)

var (
	mu          sync.Mutex
	stopGuard   chan struct{}
	backupPath  = os.Getenv("APPDATA") + `\focus\hosts.backup`
)

func Block(sites []string) error {
	mu.Lock()
	defer mu.Unlock()

	original, err := os.ReadFile(hostsPath)
	if err != nil {
		return fmt.Errorf("cannot read hosts file: %w", err)
	}

	if err := os.WriteFile(backupPath, original, 0644); err != nil {
		return fmt.Errorf("cannot write backup: %w", err)
	}

	patched := buildHosts(string(original), sites)
	if err := os.WriteFile(hostsPath, []byte(patched), 0644); err != nil {
		return fmt.Errorf("cannot write hosts file: %w", err)
	}

	flushDNS()
	startGuardian(sites)
	return nil
}

func Restore() error {
	mu.Lock()
	defer mu.Unlock()

	stopGuardian()

	backup, err := os.ReadFile(backupPath)
	if err != nil {
		stripped, err2 := stripBlock()
		if err2 != nil {
			return fmt.Errorf("no backup and cannot strip: %w", err2)
		}
		_ = os.WriteFile(hostsPath, []byte(stripped), 0644)
		flushDNS()
		return nil
	}

	if err := os.WriteFile(hostsPath, backup, 0644); err != nil {
		return fmt.Errorf("cannot restore hosts: %w", err)
	}

	_ = os.Remove(backupPath)
	flushDNS()
	return nil
}

func buildHosts(original string, sites []string) string {
	stripped := removeBlock(original)
	var sb strings.Builder
	sb.WriteString(strings.TrimRight(stripped, "\r\n"))
	sb.WriteString("\n\n")
	sb.WriteString(blockStart + "\n")
	for _, site := range sites {
		domain := strings.TrimPrefix(site, "www.")
		sb.WriteString(fmt.Sprintf("127.0.0.1 %s\n", domain))
		sb.WriteString(fmt.Sprintf("127.0.0.1 www.%s\n", domain))
	}
	sb.WriteString(blockEnd + "\n")
	return sb.String()
}

func removeBlock(content string) string {
	lines := strings.Split(content, "\n")
	var out []string
	inBlock := false
	for _, line := range lines {
		if strings.TrimSpace(line) == blockStart {
			inBlock = true
			continue
		}
		if strings.TrimSpace(line) == blockEnd {
			inBlock = false
			continue
		}
		if !inBlock {
			out = append(out, line)
		}
	}
	return strings.Join(out, "\n")
}

func stripBlock() (string, error) {
	content, err := os.ReadFile(hostsPath)
	if err != nil {
		return "", err
	}
	return removeBlock(string(content)), nil
}

func startGuardian(sites []string) {
	stopGuard = make(chan struct{})
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				mu.Lock()
				content, err := os.ReadFile(hostsPath)
				if err == nil {
					if !strings.Contains(string(content), blockStart) {
						patched := buildHosts(string(content), sites)
						_ = os.WriteFile(hostsPath, []byte(patched), 0644)
						flushDNS()
					}
				}
				mu.Unlock()
			case <-stopGuard:
				return
			}
		}
	}()
}

func stopGuardian() {
	if stopGuard != nil {
		close(stopGuard)
		stopGuard = nil
	}
}

func flushDNS() {
	// Runs ipconfig /flushdns silently
	// Actual exec call done via os/exec in the real build
	execFlushDNS()
}
