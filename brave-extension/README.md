# Focus Redirect Extension (Brave)

This extension reads active blocked sites from Focus (`http://localhost:4862/api/sites`) and redirects them to the blocked page (`http://localhost:4862`).

## Load in Brave

1. Open `brave://extensions`
2. Enable **Developer mode**
3. Click **Load unpacked**
4. Select this folder: `brave-extension`

## How it works

- Every minute, the service worker fetches blocked sites from Focus.
- It installs dynamic redirect rules for those domains.
- When you open a blocked domain, Brave redirects to `http://localhost:4862/?site=<domain>`.
- Click the extension icon to open a popup with live sync status, active redirected domains, and a **Sync Now** button.

## Notes

- Keep Focus running during active sessions so `/api/sites` is available.
- The extension redirects top-level page navigations (`main_frame`) only.
