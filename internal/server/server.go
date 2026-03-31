package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"focus/internal/blocker"
	"focus/internal/session"
)

const port = 4862

var srv *http.Server

func Start(endTime time.Time) error {
	if srv != nil {
		return nil
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleBlocked(endTime))
	mux.HandleFunc("/api/status", handleStatus)
	mux.HandleFunc("/api/stop-request", handleStopRequest)
	mux.HandleFunc("/api/stop-confirm", handleStopConfirm)
	mux.HandleFunc("/api/cancel-stop", handleCancelStop)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to start local blocker server on port %d: %w", port, err)
	}

	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		_ = srv.Serve(ln)
	}()

	return nil
}

func Stop() {
	if srv != nil {
		_ = srv.Close()
		srv = nil
	}
}

func handleBlocked(endTime time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		if host == "" {
			host = "this site"
		}
		host = strings.Split(host, ":")[0]

		html := strings.ReplaceAll(blockedHTML, "{{END_TIME_MS}}", fmt.Sprintf("%d", endTime.UnixMilli()))
		html = strings.ReplaceAll(html, "{{SITE}}", host)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(html))
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	s, err := session.Load()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{"active": false})
		return
	}
	rem := s.Remaining()
	json.NewEncoder(w).Encode(map[string]any{
		"active":    rem > 0,
		"remaining": int(rem.Seconds()),
		"stopping":  s.Stopping,
		"stop_at":   s.StopAt.UnixMilli(),
	})
}

var cooldownActive = false
var cooldownEnd time.Time

func handleStopRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

func handleStopConfirm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if !cooldownActive {
		cooldownActive = true
		cooldownEnd = time.Now().Add(2 * time.Minute)

		go func() {
			time.Sleep(2 * time.Minute)
			cooldownActive = false
			session.Clear()
			blocker.Restore()
			Stop()
		}()
	}

	json.NewEncoder(w).Encode(map[string]any{
		"ok":          true,
		"cooldown_ms": cooldownEnd.UnixMilli(),
	})
}

func handleCancelStop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cooldownActive = false
	json.NewEncoder(w).Encode(map[string]any{"ok": true})
}
