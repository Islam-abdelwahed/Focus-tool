# Focus

A minimal, beautiful website blocker for Windows. No extensions, no runtime dependencies — just a single `.exe`.

---

## How it works

- Opens a clean dashboard window (WebView2 — pre-installed on all Windows 10/11 machines)
- Edits your system `hosts` file to redirect blocked sites to `localhost`
- Runs a local HTTP server on port 80 that serves a live countdown timer
- Re-writes the hosts file every 30 seconds so manual edits don't stick
- Restores everything automatically when the timer ends

---

## Requirements

| Tool | Version | Download |
|------|---------|----------|
| Go | 1.21+ | https://go.dev/dl/ |
| Windows | 10 / 11 | — |
| WebView2 Runtime | any | Pre-installed on Win 10/11. If missing: https://developer.microsoft.com/microsoft-edge/webview2/ |

---

## Build

**Option A — double-click (simplest):**

```
build.bat
```

**Option B — command line:**

```bat
go mod tidy
go build -ldflags="-H windowsgui -s -w" -o dist\focus.exe .\cmd\focus
```

Output: `dist\focus.exe` (~6–8 MB, no installer needed)

---

## Run

Right-click `focus.exe` → **Run as administrator**

The app needs admin rights to edit the system hosts file. On first launch it will auto-elevate via UAC prompt.

---

## Project structure

```
focus/
├── cmd/
│   └── focus/
│       └── main.go              # Entry point
├── internal/
│   ├── blocker/
│   │   ├── blocker.go           # Hosts file editing + 30s guardian loop
│   │   ├── dns_windows.go       # ipconfig /flushdns
│   │   └── syscall_windows.go   # Hidden console window helper
│   ├── server/
│   │   ├── server.go            # HTTP server + stop API
│   │   └── blocked_html.go      # Blocked page HTML (embedded)
│   ├── session/
│   │   ├── session.go           # State load/save/clear
│   │   └── recovery.go          # Crash recovery + auto-stop scheduler
│   ├── tray/
│   │   └── tray.go              # System tray icon + menu
│   └── ui/
│       ├── dashboard.go         # WebView2 window + JS bridge
│       ├── dashboard_html.go    # Dashboard HTML/CSS/JS (embedded)
│       └── admin_windows.go     # UAC elevation helper
├── assets/
│   └── README.txt               # Put focus.ico here
├── focus.manifest               # UAC manifest
├── focus.rc                     # Windows resource file
├── go.mod
├── build.bat                    # One-click build for Windows
└── Makefile                     # For make users
```

---

## Anti-cheat system

Stopping early requires three steps:

1. **Friction prompt** — click "I need to stop early" (small, muted link)
2. **Type STOP** — must type the word exactly to proceed
3. **2-minute cooldown** — a visible countdown before hosts are restored

Additionally, the hosts file is re-written every 30 seconds by a background goroutine, so manually editing it during a session doesn't stick.

---

## Data stored

All data is in `%APPDATA%\focus\`:

| File | Contents |
|------|----------|
| `session.json` | Active session state (end time, sites, duration) |
| `sites.json` | Your custom blocked sites list |
| `hosts.backup` | Backup of your original hosts file (only exists during a session) |

---

## Adding an icon

Place a `focus.ico` file (multi-size: 16×16, 32×32, 48×48) in the `assets/` folder.

Then install `go-winres` and rebuild:

```bat
go install github.com/tc-hib/go-winres@latest
go-winres make --in focus.rc --out cmd\focus\rsrc.syso
go build -ldflags="-H windowsgui -s -w" -o dist\focus.exe .\cmd\focus
```

Without an icon the app builds and runs fine — Windows will use a default icon.

---

## Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/webview/webview_go` | Native WebView2 window for the dashboard UI |
| `github.com/getlantern/systray` | System tray icon with remaining time |

Both are pulled automatically by `go mod tidy`.

---

## License

MIT
