package session

import (
	"os"
	"time"
)

// RecoverIfNeeded checks on startup whether a previous session crashed
// without restoring the hosts file, and cleans up if so.
// Call this at the very start of main(), before doing anything else.
func RecoverIfNeeded() {
	backup := backupHostsPath()
	stateExists := fileExists(stateFile)
	backupExists := fileExists(backup)

	if !backupExists && !stateExists {
		return // clean state, nothing to do
	}

	s, err := Load()

	// If there's a backup but no session, or session has expired → restore
	if backupExists && (err != nil || s.Remaining() <= 0) {
		restoreHostsFromBackup(backup)
		Clear()
		return
	}

	// If session is still valid, let normal startup handle it
	if err == nil && s.Remaining() > 0 {
		return
	}

	// Fallback: remove stale state
	Clear()
}

func backupHostsPath() string {
	return os.Getenv("APPDATA") + `\focus\hosts.backup`
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func restoreHostsFromBackup(backupPath string) {
	content, err := os.ReadFile(backupPath)
	if err != nil {
		return
	}
	_ = os.WriteFile(hostsPath, content, 0644)
	_ = os.Remove(backupPath)
}

// ResumeIfActive re-attaches the guardian loop and HTTP server if a valid
// session is found on startup (e.g. user restarted the app mid-session).
func ResumeIfActive() (*State, bool) {
	s, err := Load()
	if err != nil {
		return nil, false
	}
	if s.Remaining() <= 0 {
		Clear()
		return nil, false
	}
	return s, true
}

const hostsPath = `C:\Windows\System32\drivers\etc\hosts`

// ScheduleAutoStop sets a goroutine to fire when the session timer expires.
// Safe to call multiple times — only the most recent goroutine will matter
// since session.Clear() is idempotent.
func ScheduleAutoStop(s *State, onStop func()) {
	go func() {
		rem := s.Remaining()
		if rem <= 0 {
			onStop()
			return
		}
		timer := time.NewTimer(rem)
		<-timer.C
		onStop()
	}()
}
