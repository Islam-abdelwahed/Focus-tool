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
    --bg-base: #f0f2f5;
    --bg-app: #ffffff;
    --bg-glass: rgba(255, 255, 255, 0.75);
    --bg-glass-hover: rgba(255, 255, 255, 0.95);
    --border: rgba(0, 0, 0, 0.06);
    --border2: rgba(0, 0, 0, 0.12);
    --text: #0f172a;
    --muted: #64748b;
    --hint: #94a3b8;
    --accent: #2563eb;
    --accent-hover: #1d4ed8;
    --accent-text: #ffffff;
    --success: #059669;
    --success-bg: #d1fae5;
    --success-border: #a7f3d0;
    --danger: #e11d48;
    --danger-bg: #ffe4e6;
    --danger-border: #fecdd3;
    --radius-lg: 16px;
    --radius: 12px;
    --radius-sm: 8px;
    --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -2px rgba(0, 0, 0, 0.05);
    --shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -4px rgba(0, 0, 0, 0.05);
  }

  body {
    font-family: -apple-system, BlinkMacSystemFont, "Inter", "Segoe UI", system-ui, sans-serif;
    background: var(--bg-base);
    background-image: radial-gradient(at 0% 0%, hsla(253,16%,7deg,0.03) 0, transparent 50%), radial-gradient(at 50% 0%, hsla(225,39%,30%,0.03) 0, transparent 50%), radial-gradient(at 100% 0%, hsla(339,49%,30%,0.03) 0, transparent 50%);
    color: var(--text);
    height: 100vh;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    user-select: none;
    -webkit-font-smoothing: antialiased;
  }

  .titlebar {
    background: transparent;
    padding: 12px 16px;
    display: flex;
    align-items: center;
    gap: 8px;
    -webkit-app-region: drag;
    flex-shrink: 0;
  }

  .dot { width: 12px; height: 12px; border-radius: 50%; box-shadow: inset 0 0 0 1px rgba(0,0,0,0.1); }
  .titlebar-name {
    font-size: 13px;
    font-weight: 500;
    color: var(--muted);
    margin: 0 auto;
    letter-spacing: .02em;
  }

  .body {
    flex: 1;
    overflow-y: auto;
    padding: 16px 28px 40px;
    background: var(--bg-app);
    margin: 0 12px 12px;
    border-radius: var(--radius-lg);
    box-shadow: var(--shadow-sm);
    border: 1px solid var(--border);
    display: flex;
    flex-direction: column;
  }
  
  .body::-webkit-scrollbar {
    width: 6px;
  }
  .body::-webkit-scrollbar-track {
    background: transparent;
  }
  .body::-webkit-scrollbar-thumb {
    background: var(--border2);
    border-radius: 10px;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 14px;
    margin-bottom: 32px;
    margin-top: 8px;
  }

  .logo-mark {
    width: 44px; height: 44px;
    background: linear-gradient(135deg, var(--accent), #4f46e5);
    border-radius: 12px;
    display: flex; align-items: center; justify-content: center;
    flex-shrink: 0;
    box-shadow: 0 4px 12px rgba(37, 99, 235, 0.3);
  }

  .logo-mark svg { display: block; width: 22px; height: 22px; }
  .logo-name { font-size: 22px; font-weight: 700; letter-spacing: -0.02em; color: var(--text); }
  .logo-tag { font-size: 13px; color: var(--muted); margin-top: 2px; font-weight: 400; }

  .extension-card {
    background: linear-gradient(145deg, #0f172a, #172554);
    border: 1px solid rgba(191, 219, 254, 0.3);
    border-radius: var(--radius);
    padding: 18px;
    margin-bottom: 22px;
    color: #e2e8f0;
    box-shadow: 0 16px 30px -16px rgba(15, 23, 42, 0.7);
  }

  .extension-card-head {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    gap: 16px;
    margin-bottom: 8px;
  }

  .extension-card-title {
    font-size: 16px;
    font-weight: 700;
    color: #eff6ff;
    margin-bottom: 4px;
    letter-spacing: 0.01em;
  }

  .extension-card-sub {
    font-size: 13px;
    line-height: 1.5;
    color: #bfdbfe;
  }

  .extension-pill {
    font-size: 11px;
    text-transform: uppercase;
    letter-spacing: 0.07em;
    font-weight: 700;
    color: #dbeafe;
    border: 1px solid rgba(147, 197, 253, 0.4);
    border-radius: 999px;
    padding: 5px 10px;
    background: rgba(30, 64, 175, 0.35);
    white-space: nowrap;
  }

  .extension-actions {
    display: flex;
    gap: 8px;
    margin-top: 14px;
    flex-wrap: wrap;
  }

  .extension-btn {
    font-size: 12px;
    font-weight: 600;
    border-radius: 999px;
    border: 1px solid rgba(147, 197, 253, 0.45);
    background: rgba(30, 64, 175, 0.28);
    color: #dbeafe;
    padding: 7px 12px;
    cursor: pointer;
    transition: all .2s;
  }

  .extension-btn:hover { background: rgba(59, 130, 246, 0.35); }

  .extension-btn.primary {
    background: linear-gradient(135deg, #38bdf8, #60a5fa);
    border-color: rgba(191, 219, 254, 0.8);
    color: #082f49;
  }

  .extension-btn.primary:hover {
    filter: brightness(1.05);
  }

  .extension-help {
    display: none;
    margin-top: 12px;
    background: rgba(15, 23, 42, 0.55);
    border: 1px solid rgba(147, 197, 253, 0.25);
    border-radius: 10px;
    padding: 12px;
    font-size: 12px;
    color: #cbd5e1;
    line-height: 1.65;
  }

  .extension-help strong {
    color: #eff6ff;
    font-weight: 600;
  }

  .section-label {
    font-size: 11px;
    font-weight: 600;
    letter-spacing: .08em;
    text-transform: uppercase;
    color: var(--hint);
    margin-bottom: 12px;
  }

  .pills {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
    margin-bottom: 16px;
  }

  .pill {
    padding: 8px 16px;
    border-radius: 999px;
    border: 1px solid var(--border2);
    font-size: 13px;
    font-weight: 500;
    color: var(--text);
    cursor: pointer;
    background: var(--bg-app);
    font-family: inherit;
    transition: all .2s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .pill:hover { background: var(--bg-base); border-color: var(--hint); }
  .pill.active {
    background: var(--text);
    color: var(--bg-app);
    border-color: var(--text);
    box-shadow: var(--shadow-sm);
  }

  .custom-row {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 28px;
    background: var(--bg-base);
    padding: 6px 6px 6px 16px;
    border-radius: 999px;
    width: max-content;
    border: 1px solid var(--border);
  }
  
  .custom-row span { font-size: 13px; color: var(--muted); font-weight: 500; }

  .custom-row input {
    width: 60px;
    font-size: 14px;
    font-weight: 500;
    padding: 6px 12px;
    border-radius: 999px;
    border: 1px solid var(--border);
    background: var(--bg-app);
    color: var(--text);
    font-family: inherit;
    text-align: center;
    outline: none;
    transition: all .2s;
  }

  .custom-row input:focus { border-color: var(--accent); box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.1); }

  .start-btn {
    width: 100%;
    padding: 14px;
    border-radius: var(--radius);
    background: var(--text);
    color: var(--bg-app);
    border: none;
    font-size: 15px;
    font-weight: 600;
    cursor: pointer;
    font-family: inherit;
    letter-spacing: .01em;
    transition: all .2s;
    margin-bottom: 24px;
    box-shadow: var(--shadow-md);
  }

  .start-btn:hover:not(:disabled) { transform: translateY(-1px); box-shadow: var(--shadow-lg); }
  .start-btn:active:not(:disabled) { transform: translateY(0); box-shadow: var(--shadow-sm); }
  .start-btn:disabled { opacity: .6; cursor: not-allowed; transform: none; box-shadow: none; }

  .divider {
    border: none;
    border-top: 1px solid var(--border);
    margin: 8px 0 24px;
  }

  .sites-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 16px;
  }

  .add-btn {
    font-size: 13px;
    font-weight: 500;
    color: var(--text);
    border: 1px solid var(--border2);
    background: var(--bg-app);
    padding: 6px 12px;
    border-radius: 999px;
    cursor: pointer;
    font-family: inherit;
    transition: all .2s;
    box-shadow: var(--shadow-sm);
  }

  .add-btn:hover { background: var(--bg-base); border-color: var(--hint); }

  .site-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px 12px;
    margin-bottom: 6px;
    border-radius: var(--radius-sm);
    background: rgba(0,0,0,0.02);
    border: 1px solid transparent;
    transition: all .2s;
  }
  
  .site-row:hover {
    background: var(--bg-base);
    border-color: var(--border);
  }

  .site-name { font-size: 13px; font-weight: 500; color: var(--text); }

  .site-remove {
    background: none;
    border: none;
    font-size: 18px;
    line-height: 1;
    color: var(--hint);
    cursor: pointer;
    padding: 2px 6px;
    border-radius: 6px;
    transition: all .2s;
    font-family: inherit;
  }

  .site-remove:hover { color: var(--danger); background: var(--danger-bg); }

  .add-row {
    display: none;
    gap: 8px;
    margin-top: 12px;
    animation: fadeIn 0.2s ease;
  }
  
  @keyframes fadeIn { from { opacity: 0; transform: translateY(-4px); } to { opacity: 1; transform: translateY(0); } }

  .add-row input {
    flex: 1;
    font-size: 13px;
    padding: 10px 14px;
    border-radius: var(--radius-sm);
    border: 1px solid var(--border2);
    background: var(--bg-app);
    color: var(--text);
    font-family: inherit;
    outline: none;
    transition: all .2s;
    box-shadow: inset 0 1px 2px rgba(0,0,0,0.03);
  }

  .add-row input:focus { border-color: var(--accent); box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.1); }

  .add-row button {
    font-size: 13px;
    font-weight: 500;
    padding: 10px 18px;
    border-radius: var(--radius-sm);
    background: var(--text);
    color: var(--bg-app);
    border: none;
    cursor: pointer;
    font-family: inherit;
    transition: all .2s;
  }
  
  .add-row button:hover { opacity: 0.9; }

  .active-banner {
    background: var(--success-bg);
    border: 1px solid var(--success-border);
    border-radius: var(--radius);
    padding: 18px 20px;
    margin-bottom: 24px;
    display: none;
    box-shadow: 0 4px 12px rgba(5, 150, 105, 0.1);
  }

  .active-banner-title {
    font-size: 14px;
    font-weight: 600;
    color: var(--success);
    margin-bottom: 4px;
    display: flex;
    align-items: center;
    gap: 8px;
  }
  
  .active-banner-title::before {
    content: '';
    display: block;
    width: 8px; height: 8px;
    background: var(--success);
    border-radius: 50%;
    box-shadow: 0 0 8px var(--success);
  }

  .active-banner-sub {
    font-size: 13px;
    color: var(--success);
    opacity: .8;
    margin-left: 16px;
    font-variant-numeric: tabular-nums;
  }

  .stop-early-btn {
    width: 100%;
    margin-top: 16px;
    padding: 10px;
    border-radius: var(--radius-sm);
    background: var(--bg-glass);
    border: 1px solid var(--danger-border);
    color: var(--danger);
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
    font-family: inherit;
    transition: all .2s;
  }

  .stop-early-btn:hover { background: var(--danger-bg); border-color: rgba(225, 29, 72, 0.3); }

  .statusbar {
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    padding: 12px 16px;
    background: transparent;
  }

  .status-dot { width: 8px; height: 8px; border-radius: 50%; background: var(--hint); transition: background .3s; }
  .status-dot.on { background: var(--success); box-shadow: 0 0 8px rgba(5, 150, 105, 0.4); }
  .statusbar-text { font-size: 12px; font-weight: 500; color: var(--muted); }

  .err { font-size: 13px; font-weight: 500; color: var(--danger); margin-top: 12px; display: none; background: var(--danger-bg); padding: 10px; border-radius: var(--radius-sm); border: 1px solid var(--danger-border); }

  .modal {
    position: fixed;
    inset: 0;
    background: rgba(15, 23, 42, 0.4);
    backdrop-filter: blur(4px);
    display: none;
    align-items: center;
    justify-content: center;
    padding: 16px;
    z-index: 50;
    opacity: 0;
    transition: opacity 0.2s;
  }
  
  .modal.show {
    display: flex;
    opacity: 1;
  }

  .modal-card {
    width: 100%;
    max-width: 360px;
    background: var(--bg-app);
    border: 1px solid var(--border2);
    border-radius: var(--radius-lg);
    padding: 24px;
    box-shadow: var(--shadow-lg), 0 20px 40px rgba(0,0,0,0.1);
    transform: translateY(10px) scale(0.95);
    transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  }
  
  .modal.show .modal-card {
    transform: translateY(0) scale(1);
  }

  .modal-title {
    font-size: 16px;
    font-weight: 700;
    color: var(--text);
    margin-bottom: 8px;
  }

  .modal-text {
    font-size: 13px;
    color: var(--muted);
    line-height: 1.6;
    margin-bottom: 24px;
  }

  .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
  }

  .modal-btn {
    font-size: 13px;
    font-weight: 600;
    border-radius: var(--radius-sm);
    border: 1px solid var(--border2);
    background: var(--bg-app);
    color: var(--text);
    padding: 8px 16px;
    cursor: pointer;
    font-family: inherit;
    transition: all .2s;
  }
  
  .modal-btn:hover {
    background: var(--bg-base);
  }

  .modal-btn.danger {
    background: var(--danger);
    border-color: var(--danger);
    color: #fff;
    box-shadow: 0 2px 8px rgba(225, 29, 72, 0.25);
  }
  
  .modal-btn.danger:hover {
    background: #be123c;
    border-color: #be123c;
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

  <div class="extension-card" id="extension-card">
    <div class="extension-card-head">
      <div>
        <div class="extension-card-title">Browser extension needed for full blocking</div>
        <div class="extension-card-sub">Install the Focus Redirect extension in Brave/Chrome so blocked websites always route to this localhost page.</div>
      </div>
      <div class="extension-pill">Recommended</div>
    </div>
    <div class="extension-actions">
      <button class="extension-btn primary" onclick="toggleExtensionHelp()">Install steps</button>
      <button class="extension-btn" onclick="dismissExtensionCard()">I've installed it</button>
    </div>
    <div class="extension-help" id="extension-help">
      <strong>1)</strong> Open <strong>brave://extensions</strong> in your browser.<br>
      <strong>2)</strong> Enable <strong>Developer mode</strong>.<br>
      <strong>3)</strong> Click <strong>Load unpacked</strong>.<br>
      <strong>4)</strong> Select the <strong>brave-extension</strong> folder from this Focus project.
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
const EXTENSION_CARD_KEY = 'focus.extensionInstalled';

async function init() {
  initExtensionCard();

  const sitesJSON = await bridge_getSites();
  sites = JSON.parse(sitesJSON);
  renderSites();

  const statusJSON = await bridge_getStatus();
  const status = JSON.parse(statusJSON);
  if (status.active) {
    setActiveState(status.end_ms);
  }
}

function initExtensionCard() {
  try {
    if (localStorage.getItem(EXTENSION_CARD_KEY) === '1') {
      const card = document.getElementById('extension-card');
      if (card) card.style.display = 'none';
    }
  } catch (_e) {}
}

function toggleExtensionHelp() {
  const help = document.getElementById('extension-help');
  if (!help) return;
  help.style.display = help.style.display === 'block' ? 'none' : 'block';
}

function dismissExtensionCard() {
  const card = document.getElementById('extension-card');
  if (card) card.style.display = 'none';
  try {
    localStorage.setItem(EXTENSION_CARD_KEY, '1');
  } catch (_e) {}
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
  modal.classList.add('show');

  return new Promise(resolve => {
    stopModalResolver = resolve;
  });
}

function closeStopModal(confirmed) {
  const modal = document.getElementById('stop-modal');
  modal.classList.remove('show');
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
