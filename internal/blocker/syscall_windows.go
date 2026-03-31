//go:build windows

package blocker

import "syscall"

func hiddenWindow() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{HideWindow: true}
}
