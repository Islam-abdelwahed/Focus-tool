//go:build windows

package blocker

import (
	"os/exec"
)

func execFlushDNS() {
	cmd := exec.Command("ipconfig", "/flushdns")
	cmd.SysProcAttr = hiddenWindow()
	_ = cmd.Run()
}
