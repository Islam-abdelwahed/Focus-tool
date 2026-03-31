package tray

import (
	"fmt"
	"time"

	"fyne.io/systray"
)

func Run(endTime time.Time, onStop func()) {
	go systray.Run(func() {
		systray.SetTitle("Focus")
		systray.SetTooltip("Focus session active")
		systray.SetIcon(minimalIcon())

		mStatus := systray.AddMenuItem("Starting…", "")
		mStatus.Disable()
		systray.AddSeparator()
		mStop := systray.AddMenuItem("Stop session", "")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "")

		go func() {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					rem := time.Until(endTime)
					if rem <= 0 {
						mStatus.SetTitle("Session complete")
						systray.Quit()
						return
					}
					h := int(rem.Hours())
					m := int(rem.Minutes()) % 60
					s := int(rem.Seconds()) % 60
					var label string
					if h > 0 {
						label = fmt.Sprintf("Focus · %dh %02dm remaining", h, m)
					} else {
						label = fmt.Sprintf("Focus · %dm %02ds remaining", m, s)
					}
					mStatus.SetTitle(label)
					if h > 0 {
						systray.SetTitle(fmt.Sprintf("%dh%02dm", h, m))
					} else {
						systray.SetTitle(fmt.Sprintf("%d:%02d", m, s))
					}

				case <-mStop.ClickedCh:
					onStop()
					systray.Quit()
					return

				case <-mQuit.ClickedCh:
					systray.Quit()
					return
				}
			}
		}()

	}, func() {})
}

// minimalIcon returns a minimal valid ICO byte slice (16x16 black square).
// Replace with a real icon by embedding focus.ico via go:embed.
func minimalIcon() []byte {
	return []byte{
		0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x10, 0x10,
		0x00, 0x00, 0x01, 0x00, 0x04, 0x00, 0x28, 0x01,
		0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x28, 0x00,
		0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x20, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
}
