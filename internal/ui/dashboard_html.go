package ui

const dashboardHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Focus</title>
<style>
  *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

  :root {
    --bg: #ffffff;
    --bg2: #f7f7f5;
    --bg3: #f0f0ed;
    --border: rgba(0,0,0,0.09);
    --border2: rgba(0,0,0,0.15);
    --text: #1a1a18;
    --muted: #6b6b67;
    --hint: #a8a8a4;
    --accent: #1a1a18;
    --accent-text: #ffffff;
    --success: #3b6d11;
    --success-bg: #eaf3de;
    --danger: #a32d2d;
    --radius: 10px;
    --radius-sm: 7px;
  }

  body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", system-ui, sans-serif;
    background: var(--bg);
    color: var(--text);
    height: 100vh;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    user-select: none;
  }

  .titlebar {
    background: var(--bg2);
    border-bottom: 0.5px solid var(--border);
    padding: 10px 16px;
    display: flex;
    align-items: center;
    gap: 8px;
    -webkit-app-region: drag;
    flex-shrink: 0;
  }

  .dot { width: 12px; height: 12px; border-radius: 50%; }
  .titlebar-name {
    font-size: 13px;
    color: var(--muted);
    margin: 0 auto;
    letter-spacing: .01em;
  }

  .body {
    flex: 1;
    overflow-y: auto;
    padding: 24px 24px 32px;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 28px;
  }

  .logo-mark {
    width: 32px; height: 32px;
    background: var(--accent);
    border-radius: 8px;
    display: flex; align-items: center; justify-content: center;
    flex-shrink: 0;
  }

  .logo-mark svg { display: block; }
  .logo-name { font-size: 17px; font-weight: 500; }
  .logo-tag { font-size: 12px; color: var(--hint); margin-top: 1px; }

  .section-label {
    font-size: 10px;
    font-weight: 500;
    letter-spacing: .08em;
    text-transform: uppercase;
    color: var(--hint);
    margin-bottom: 10px;
  }

  .pills {
    display: flex;
    gap: 6px;
    flex-wrap: wrap;
    margin-bottom: 14px;
  }

  .pill {
    padding: 6px 14px;
    border-radius: 999px;
    border: 0.5px solid var(--border2);
    font-size: 13px;
    color: var(--muted);
    cursor: pointer;
    background: transparent;
    font-family: inherit;
    transition: all .12s;
  }

  .pill:hover { background: var(--bg2); }
  .pill.active {
    background: var(--accent);
    color: var(--accent-text);
    border-color: var(--accent);
  }

  .custom-row {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 22px;
  }

  .custom-row input {
    width: 70px;
    font-size: 14px;
    padding: 6px 10px;
    border-radius: var(--radius-sm);
    border: 0.5px solid var(--border2);
    background: var(--bg);
    color: var(--text);
    font-family: inherit;
    text-align: center;
    outline: none;
  }

  .custom-row input:focus { border-color: var(--accent); }
  .custom-row span { font-size: 13px; color: var(--muted); }

  .start-btn {
    width: 100%;
    padding: 12px;
    border-radius: var(--radius);
    background: var(--accent);
    color: var(--accent-text);
    border: none;
    font-size: 14px;
    font-weight: 500;
    cursor: pointer;
    font-family: inherit;
    letter-spacing: .01em;
    transition: opacity .15s;
    margin-bottom: 22px;
  }

  .start-btn:hover:not(:disabled) { opacity: .85; }
  .start-btn:disabled { opacity: .5; cursor: not-allowed; }
  .start-btn.active-state {
    background: var(--success-bg);
    color: var(--success);
  }

  .divider {
    border: none;
    border-top: 0.5px solid var(--border);
    margin-bottom: 20px;
  }

  .sites-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 12px;
  }

  .add-btn {
    font-size: 12px;
    color: var(--muted);
    border: 0.5px solid var(--border2);
    background: transparent;
    padding: 4px 10px;
    border-radius: var(--radius-sm);
    cursor: pointer;
    font-family: inherit;
    transition: background .12s;
  }

  .add-btn:hover { background: var(--bg2); }

  .site-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 8px 0;
    border-bottom: 0.5px solid var(--border);
  }

  .site-row:last-child { border-bottom: none; }
  .site-name { font-size: 13px; }

  .site-remove {
    background: none;
    border: none;
    font-size: 18px;
    line-height: 1;
    color: var(--hint);
    cursor: pointer;
    padding: 0 2px;
    transition: color .12s;
    font-family: inherit;
  }

  .site-remove:hover { color: var(--danger); }

  .add-row {
    display: none;
    gap: 8px;
    margin-top: 12px;
  }

  .add-row input {
    flex: 1;
    font-size: 13px;
    padding: 7px 10px;
    border-radius: var(--radius-sm);
    border: 0.5px solid var(--border2);
    background: var(--bg);
    color: var(--text);
    font-family: inherit;
    outline: none;
  }

  .add-row input:focus { border-color: var(--accent); }

  .add-row button {
    font-size: 13px;
    padding: 7px 14px;
    border-radius: var(--radius-sm);
    background: var(--accent);
    color: var(--accent-text);
    border: none;
    cursor: pointer;
    font-family: inherit;
  }

  .active-banner {
    background: var(--success-bg);
    border: 0.5px solid rgba(59,109,17,0.2);
    border-radius: var(--radius);
    padding: 14px 16px;
    margin-bottom: 20px;
    display: none;
  }

  .active-banner-title {
    font-size: 13px;
    font-weight: 500;
    color: var(--success);
    margin-bottom: 2px;
  }

  .active-banner-sub {
    font-size: 12px;
    color: var(--success);
    opacity: .7;
  }

  .stop-early-btn {
    width: 100%;
    margin-top: 10px;
    padding: 9px;
    border-radius: var(--radius-sm);
    background: transparent;
    border: 0.5px solid rgba(163,45,45,0.35);
    color: var(--danger);
    font-size: 13px;
    cursor: pointer;
    font-family: inherit;
    transition: background .12s;
  }

  .stop-early-btn:hover { background: rgba(163,45,45,0.06); }

  .statusbar {
    flex-shrink: 0;
    display: flex;
    align-items: center;
    gap: 6px;
    padding: 9px 16px;
    background: var(--bg2);
    border-top: 0.5px solid var(--border);
  }

  .status-dot { width: 7px; height: 7px; border-radius: 50%; background: var(--hint); }
  .status-dot.on { background: #639922; }
  .statusbar-text { font-size: 12px; color: var(--muted); }

  .err { font-size: 12px; color: var(--danger); margin-top: 8px; display: none; }

  .modal {
    position: fixed;
    inset: 0;
    background: rgba(0,0,0,0.35);
    display: none;
    align-items: center;
    justify-content: center;
    padding: 16px;
    z-index: 10;
  }

  .modal-card {
    width: 100%;
    max-width: 360px;
    background: var(--bg);
    border: 0.5px solid var(--border2);
    border-radius: var(--radius);
    padding: 16px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.12);
  }

  .modal-title {
    font-size: 14px;
    font-weight: 600;
    margin-bottom: 6px;
  }

  .modal-text {
    font-size: 12px;
    color: var(--muted);
    line-height: 1.5;
    margin-bottom: 14px;
  }

  .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 8px;
  }

  .modal-btn {
    font-size: 12px;
    border-radius: var(--radius-sm);
    border: 0.5px solid var(--border2);
    background: transparent;
    color: var(--text);
    padding: 7px 10px;
    cursor: pointer;
    font-family: inherit;
  }

  .modal-btn.danger {
    background: var(--danger);
    border-color: var(--danger);
    color: #fff;
  }
</style>
</head>
<body>

<div class="titlebar">
  <div class="dot" style="background:#E24B4A"></div>
  <div class="dot" style="background:#EF9F27"></div>
  <div class="dot" style="background:#639922"></div>
  <span class="titlebar-name">Focus</span>
</div>

<div class="body">

  <div class="logo">
    <div class="logo-mark">
      <svg width="18" height="18" viewBox="0 0 18 18" fill="none">
        <rect x="2.5" y="2.5" width="13" height="13" rx="3" stroke="white" stroke-width="1.4"/>
        <line x1="9" y1="5.5" x2="9" y2="12.5" stroke="white" stroke-width="1.4" stroke-linecap="round"/>
        <line x1="5.5" y1="9" x2="12.5" y2="9" stroke="white" stroke-width="1.4" stroke-linecap="round"/>
      </svg>
    </div>
    <div>
      <div class="logo-name">Focus</div>
      <div class="logo-tag" id="logo-tag">No session active</div>
    </div>
  </div>

  <div id="active-banner" class="active-banner">
    <div class="active-banner-title">Session active</div>
    <div class="active-banner-sub" id="banner-sub">--:-- remaining</div>
    <button class="stop-early-btn" onclick="stopEarly()">Stop session early</button>
  </div>

  <div id="start-section">
    <div class="section-label">Duration</div>
    <div class="pills" id="pills">
      <button class="pill" onclick="selectPill(this,25)">25 min</button>
      <button class="pill active" onclick="selectPill(this,50)">50 min</button>
      <button class="pill" onclick="selectPill(this,90)">90 min</button>
      <button class="pill" onclick="selectPill(this,120)">2 hours</button>
    </div>
    <div class="custom-row">
      <input type="number" id="custom-min" min="1" max="480" placeholder="—"
             oninput="onCustomInput(this.value)">
      <span>min custom</span>
    </div>

    <button class="start-btn" id="start-btn" onclick="startSession()">Start focus session</button>
    <div class="err" id="err-msg"></div>
  </div>

  <hr class="divider">

  <div class="sites-header">
    <span class="section-label" style="margin-bottom:0">Blocked sites</span>
    <button class="add-btn" onclick="toggleAdd()">+ Add</button>
  </div>

  <div id="site-list"></div>

  <div class="add-row" id="add-row">
    <input type="text" id="new-site" placeholder="e.g. youtube.com"
           onkeydown="if(event.key==='Enter')addSite()">
    <button onclick="addSite()">Add</button>
  </div>

</div>

<div class="statusbar">
  <div class="status-dot" id="status-dot"></div>
  <span class="statusbar-text" id="status-text">Ready</span>
</div>

<div class="modal" id="stop-modal" role="dialog" aria-modal="true" aria-labelledby="stop-modal-title">
  <div class="modal-card">
    <div class="modal-title" id="stop-modal-title">Stop session early?</div>
    <div class="modal-text">This will immediately unblock all blocked sites and end your current focus session.</div>
    <div class="modal-actions">
      <button class="modal-btn" id="stop-cancel-btn">Cancel</button>
      <button class="modal-btn danger" id="stop-confirm-btn">Stop session</button>
    </div>
  </div>
</div>

<script>
let selectedMin = 50;
let sites = [];
let sessionActive = false;
let sessionEndMs = 0;
let timerInterval = null;
let stopModalResolver = null;

async function init() {
  const sitesJSON = await bridge_getSites();
  sites = JSON.parse(sitesJSON);
  renderSites();

  const statusJSON = await bridge_getStatus();
  const status = JSON.parse(statusJSON);
  if (status.active) {
    setActiveState(status.end_ms);
  }
}

function renderSites() {
  const list = document.getElementById('site-list');
  list.innerHTML = '';
  sites.forEach((s, i) => {
    const row = document.createElement('div');
    row.className = 'site-row';
    row.innerHTML =
      '<span class="site-name">'+escHtml(s)+'</span>' +
      '<button class="site-remove" onclick="removeSite('+i+')">&#215;</button>';
    list.appendChild(row);
  });
}

function escHtml(s) {
  return s.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;');
}

async function persistSites() {
  await bridge_saveSites(JSON.stringify(sites));
}

function removeSite(i) {
  sites.splice(i, 1);
  renderSites();
  persistSites();
}

function toggleAdd() {
  const row = document.getElementById('add-row');
  const showing = row.style.display === 'flex';
  row.style.display = showing ? 'none' : 'flex';
  if (!showing) document.getElementById('new-site').focus();
}

function addSite() {
  const input = document.getElementById('new-site');
  let val = input.value.trim().replace(/^https?:\/\//, '').replace(/\/.*$/, '').toLowerCase();
  if (!val) return;
  if (!sites.includes(val)) {
    sites.push(val);
    renderSites();
    persistSites();
  }
  input.value = '';
  document.getElementById('add-row').style.display = 'none';
}

function selectPill(el, val) {
  document.querySelectorAll('.pill').forEach(p => p.classList.remove('active'));
  el.classList.add('active');
  selectedMin = val;
  document.getElementById('custom-min').value = '';
  updateStatusBar();
}

function onCustomInput(val) {
  if (val && parseInt(val) > 0) {
    selectedMin = parseInt(val);
    document.querySelectorAll('.pill').forEach(p => p.classList.remove('active'));
  }
}

function updateStatusBar() {
  const h = Math.floor(selectedMin / 60);
  const m = selectedMin % 60;
  const label = h > 0 ? (m > 0 ? h+'h '+m+'m' : h+'h') : m+' min';
  document.getElementById('status-text').textContent = 'Ready · '+label+' selected';
}

async function startSession() {
  const btn = document.getElementById('start-btn');
  const errEl = document.getElementById('err-msg');
  errEl.style.display = 'none';
  errEl.style.color = 'var(--danger)';
  btn.disabled = true;
  btn.textContent = 'Starting…';

  const result = await bridge_start(selectedMin);
  const data = JSON.parse(result);

  if (!data.ok) {
    btn.disabled = false;
    btn.textContent = 'Start focus session';
    errEl.textContent = data.error || 'Failed to start. Run as Administrator.';
    errEl.style.display = 'block';
    return;
  }

  if (data.warning) {
    errEl.textContent = data.warning;
    errEl.style.color = 'var(--muted)';
    errEl.style.display = 'block';
  }

  setActiveState(data.end_ms);
}

function setActiveState(endMs) {
  sessionActive = true;
  sessionEndMs = endMs;

  document.getElementById('active-banner').style.display = 'block';
  document.getElementById('start-section').style.display = 'none';
  document.getElementById('logo-tag').textContent = 'Session active';
  document.getElementById('status-dot').classList.add('on');

  if (timerInterval) clearInterval(timerInterval);
  timerInterval = setInterval(updateBannerTimer, 1000);
  updateBannerTimer();
}

function updateBannerTimer() {
  const rem = sessionEndMs - Date.now();
  if (rem <= 0) {
    clearInterval(timerInterval);
    resetToIdle();
    return;
  }
  const s = Math.floor(rem / 1000);
  const h = Math.floor(s / 3600);
  const m = Math.floor((s % 3600) / 60);
  const sec = s % 60;
  let label = '';
  if (h > 0) label = h+'h '+String(m).padStart(2,'0')+'m';
  else label = m+'m '+String(sec).padStart(2,'0')+'s';
  document.getElementById('banner-sub').textContent = label+' remaining';
  document.getElementById('status-text').textContent = 'Blocking '+sites.length+' sites · '+label;
}

async function stopEarly() {
  const confirmed = await showStopModal();
  if (!confirmed) return;
  await bridge_stop();
  resetToIdle();
}

function showStopModal() {
  const modal = document.getElementById('stop-modal');
  modal.style.display = 'flex';

  return new Promise(resolve => {
    stopModalResolver = resolve;
  });
}

function closeStopModal(confirmed) {
  const modal = document.getElementById('stop-modal');
  modal.style.display = 'none';
  if (stopModalResolver) {
    stopModalResolver(confirmed);
    stopModalResolver = null;
  }
}

function resetToIdle() {
  sessionActive = false;
  if (timerInterval) clearInterval(timerInterval);
  document.getElementById('active-banner').style.display = 'none';
  document.getElementById('start-section').style.display = 'block';
  document.getElementById('start-btn').disabled = false;
  document.getElementById('start-btn').textContent = 'Start focus session';
  document.getElementById('logo-tag').textContent = 'No session active';
  document.getElementById('status-dot').classList.remove('on');
  document.getElementById('status-text').textContent = 'Ready';
}

document.getElementById('stop-cancel-btn').addEventListener('click', () => closeStopModal(false));
document.getElementById('stop-confirm-btn').addEventListener('click', () => closeStopModal(true));
document.getElementById('stop-modal').addEventListener('click', (e) => {
  if (e.target.id === 'stop-modal') closeStopModal(false);
});

updateStatusBar();
init();
</script>
</body>
</html>`
