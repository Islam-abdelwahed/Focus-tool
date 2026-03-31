package ui

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	webview2 "github.com/jchv/go-webview2"

	"focus/internal/blocker"
	"focus/internal/server"
	"focus/internal/session"
)

func RunDashboard() {
	session.RecoverIfNeeded()

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     false,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Focus",
			Width:  480,
			Height: 680,
		},
	})
	if w == nil {
		panic("Failed to create webview — is WebView2 runtime installed?")
	}
	defer w.Destroy()

	w.SetSize(480, 680, webview2.HintFixed)
	bindBridge(w)
	w.SetHtml(dashboardHTML)

	// Re-attach server if session survived a restart
	if s, active := session.ResumeIfActive(); active {
		_ = server.Start(s.EndTime)
		session.ScheduleAutoStop(s, func() {
			session.Clear()
			blocker.Restore()
			server.Stop()
		})
	}

	w.Run()
}

func bindBridge(w webview2.WebView) {
	w.Bind("bridge_getSites", func() string {
		sites := session.LoadSites()
		b, _ := json.Marshal(sites)
		return string(b)
	})

	w.Bind("bridge_saveSites", func(sitesJSON string) string {
		var sites []string
		if err := json.Unmarshal([]byte(sitesJSON), &sites); err != nil {
			return `{"ok":false,"error":"invalid json"}`
		}
		if err := session.SaveSites(sites); err != nil {
			return fmt.Sprintf(`{"ok":false,"error":%q}`, err.Error())
		}
		return `{"ok":true}`
	})

	w.Bind("bridge_getStatus", func() string {
		s, err := session.Load()
		if err != nil {
			return `{"active":false}`
		}
		rem := s.Remaining()
		if rem <= 0 {
			session.Clear()
			return `{"active":false}`
		}
		b, _ := json.Marshal(map[string]any{
			"active":    true,
			"remaining": int(rem.Seconds()),
			"end_ms":    s.EndTimeMs(),
		})
		return string(b)
	})

	w.Bind("bridge_start", func(minutes int) string {
		if session.IsActive() {
			return `{"ok":false,"error":"session already active"}`
		}

		sites := session.LoadSites()
		endTime := time.Now().Add(time.Duration(minutes) * time.Minute)

		s := &session.State{
			EndTime:     endTime,
			DurationMin: minutes,
			Sites:       sites,
		}

		if err := session.Save(s); err != nil {
			return fmt.Sprintf(`{"ok":false,"error":%q}`, err.Error())
		}

		if err := blocker.Block(sites); err != nil {
			session.Clear()
			server.Stop()
			return fmt.Sprintf(`{"ok":false,"error":%q}`, err.Error())
		}

		serverWarning := ""
		if err := server.Start(endTime); err != nil {
			serverWarning = "Port 80 is unavailable, so browser timer page could not be opened. Session is still active."
		}

		session.ScheduleAutoStop(s, func() {
			session.Clear()
			blocker.Restore()
			server.Stop()
		})

		if serverWarning == "" {
			openBrowser("http://localhost:4862")
		}

		b, _ := json.Marshal(map[string]any{
			"ok":      true,
			"end_ms":  endTime.UnixMilli(),
			"warning": serverWarning,
		})
		return string(b)
	})

	w.Bind("bridge_stop", func() string {
		session.Clear()
		blocker.Restore()
		server.Stop()
		return `{"ok":true}`
	})
}

func openBrowser(url string) {
	_ = exec.Command("cmd", "/c", "start", url).Start()
}
