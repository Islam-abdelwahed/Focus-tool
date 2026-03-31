//go:build windows

package ui

import (
	"os"
	"os/exec"
	"syscall"
)

func IsAdmin() bool {
	_, err := os.Open(`\\.\PHYSICALDRIVE0`)
	return err == nil
}

func RelaunchAsAdmin() {
	exe, _ := os.Executable()
	cmd := exec.Command("cmd", "/c", "start", "", "/wait",
		"powershell", "-Command",
		"Start-Process", "-FilePath", `"`+exe+`"`,
		"-Verb", "RunAs",
	)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_ = cmd.Start()
	os.Exit(0)
}
