package server

const blockedHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Focus session active</title>
<style>
  *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

  :root {
    --bg: #ffffff;
    --bg2: #f7f7f5;
    --border: rgba(0,0,0,0.10);
    --text: #1a1a18;
    --muted: #6b6b67;
    --hint: #a8a8a4;
    --radius: 12px;
    --radius-sm: 8px;
  }

  @media (prefers-color-scheme: dark) {
    :root {
      --bg: #1c1c1a;
      --bg2: #242422;
      --border: rgba(255,255,255,0.10);
      --text: #e8e8e4;
      --muted: #9a9a96;
      --hint: #5a5a58;
    }
  }

  body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", system-ui, sans-serif;
    background: var(--bg);
    color: var(--text);
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 24px;
  }

  .card {
    width: 100%;
    max-width: 480px;
    text-align: center;
  }

  .site-badge {
    display: inline-block;
    font-size: 11px;
    font-weight: 500;
    letter-spacing: .07em;
    text-transform: uppercase;
    color: var(--hint);
    background: var(--bg2);
    border: 0.5px solid var(--border);
    border-radius: 999px;
    padding: 4px 14px;
    margin-bottom: 28px;
  }

  h1 {
    font-size: 22px;
    font-weight: 500;
    color: var(--text);
    margin-bottom: 6px;
    letter-spacing: -.01em;
  }

  .sub {
    font-size: 14px;
    color: var(--muted);
    line-height: 1.65;
    margin-bottom: 32px;
  }

  .timer-card {
    background: var(--bg2);
    border: 0.5px solid var(--border);
    border-radius: var(--radius);
    padding: 28px 24px 20px;
    margin-bottom: 28px;
  }

  .timer-label {
    font-size: 11px;
    font-weight: 500;
    letter-spacing: .08em;
    text-transform: uppercase;
    color: var(--hint);
    margin-bottom: 10px;
  }

  .timer-display {
    font-size: 48px;
    font-weight: 500;
    color: var(--text);
    letter-spacing: .04em;
    font-variant-numeric: tabular-nums;
    line-height: 1;
  }

  .timer-segments {
    display: flex;
    justify-content: center;
    gap: 0;
    margin-top: 8px;
  }

  .seg-label {
    font-size: 11px;
    color: var(--hint);
    width: 64px;
    text-align: center;
  }

  .quote-area {
    min-height: 60px;
    margin-bottom: 32px;
    padding: 0 8px;
  }

  .quote-text {
    font-size: 14px;
    color: var(--muted);
    line-height: 1.7;
    font-style: italic;
    transition: opacity .4s ease;
  }

  .quote-author {
    font-size: 12px;
    color: var(--hint);
    margin-top: 6px;
    transition: opacity .4s ease;
  }

  .stop-area {
    border-top: 0.5px solid var(--border);
    padding-top: 20px;
  }

  .stop-link {
    font-size: 12px;
    color: var(--hint);
    background: none;
    border: none;
    cursor: pointer;
    text-decoration: underline;
    text-decoration-color: var(--border);
    font-family: inherit;
    transition: color .15s;
  }

  .stop-link:hover { color: var(--muted); }

  .flow-box {
    background: var(--bg2);
    border: 0.5px solid var(--border);
    border-radius: var(--radius-sm);
    padding: 16px;
    margin-top: 14px;
    text-align: left;
    display: none;
  }

  .flow-label {
    font-size: 13px;
    color: var(--muted);
    margin-bottom: 10px;
    line-height: 1.5;
  }

  .flow-label strong {
    color: var(--text);
    font-weight: 500;
    background: var(--bg);
    border: 0.5px solid var(--border);
    border-radius: 4px;
    padding: 1px 6px;
    font-style: normal;
    font-size: 12px;
    letter-spacing: .04em;
  }

  .confirm-input {
    width: 100%;
    font-size: 14px;
    text-align: center;
    padding: 9px 12px;
    border-radius: var(--radius-sm);
    border: 0.5px solid var(--border);
    background: var(--bg);
    color: var(--text);
    font-family: inherit;
    outline: none;
    letter-spacing: .08em;
  }

  .confirm-input:focus {
    border-color: rgba(0,0,0,0.3);
  }

  @media (prefers-color-scheme: dark) {
    .confirm-input:focus { border-color: rgba(255,255,255,0.3); }
  }

  .confirm-hint {
    font-size: 11px;
    color: var(--hint);
    margin-top: 6px;
    text-align: center;
  }

  .cooldown-display {
    font-size: 28px;
    font-weight: 500;
    color: var(--text);
    font-variant-numeric: tabular-nums;
    text-align: center;
    margin-bottom: 4px;
  }

  .cooldown-sub {
    font-size: 11px;
    color: var(--hint);
    text-align: center;
  }

  .done-msg {
    font-size: 14px;
    color: var(--muted);
    text-align: center;
  }
</style>
</head>
<body>
<div class="card">
  <div class="site-badge" id="site-badge">{{SITE}}</div>
  <h1>You're in a focus session</h1>
  <p class="sub">This site is blocked until your session ends.<br>Stay with it.</p>

  <div class="timer-card">
    <div class="timer-label">Time remaining</div>
    <div class="timer-display" id="timer">--:--:--</div>
    <div class="timer-segments">
      <span class="seg-label">hrs</span>
      <span class="seg-label">min</span>
      <span class="seg-label">sec</span>
    </div>
  </div>

  <div class="quote-area">
    <div class="quote-text" id="qt"></div>
    <div class="quote-author" id="qa"></div>
  </div>

  <div class="stop-area">
    <button class="stop-link" id="stop-btn" onclick="phase1()">I need to stop early</button>

    <div class="flow-box" id="box-confirm">
      <div class="flow-label">Type <strong>STOP</strong> to confirm ending your session early</div>
      <input class="confirm-input" id="confirm-in" type="text" autocomplete="off"
             placeholder="type STOP here" oninput="checkConfirm(this.value)">
      <div class="confirm-hint">This is intentionally inconvenient.</div>
    </div>

    <div class="flow-box" id="box-cooldown">
      <div class="flow-label">Stopping in — hosts will be restored when this reaches zero</div>
      <div class="cooldown-display" id="cd-display">2:00</div>
      <div class="cooldown-sub">Close this tab if you change your mind — it won't help.</div>
    </div>

    <div class="flow-box" id="box-done">
      <div class="done-msg">Session ended. Sites are unblocked.</div>
    </div>
  </div>
</div>

<script>
const END_MS = {{END_TIME_MS}};

const quotes = [
  {t:'"The successful warrior is the average person, with laser-like focus."', a:'— Bruce Lee'},
  {t:'"It\'s not that I\'m so smart. I just stay with problems longer."', a:'— Albert Einstein'},
  {t:'"Deep work is the ability to focus without distraction on a cognitively demanding task."', a:'— Cal Newport'},
  {t:'"You don\'t have to see the whole staircase, just take the first step."', a:'— Martin Luther King Jr.'},
  {t:'"The secret of getting ahead is getting started."', a:'— Mark Twain'},
  {t:'"Energy and persistence conquer all things."', a:'— Benjamin Franklin'},
  {t:'"Concentrate all your thoughts upon the work at hand."', a:'— Alexander Graham Bell'},
  {t:'"Do the hard jobs first. The easy jobs will take care of themselves."', a:'— Dale Carnegie'},
  {t:'"The key is not to prioritize your schedule, but to schedule your priorities."', a:'— Stephen Covey'},
  {t:'"Action is the foundational key to all success."', a:'— Pablo Picasso'},
];

let qi = Math.floor(Math.random() * quotes.length);
const qt = document.getElementById('qt');
const qa = document.getElementById('qa');

function showQuote() {
  qt.textContent = quotes[qi].t;
  qa.textContent = quotes[qi].a;
}

function nextQuote() {
  qt.style.opacity = '0';
  qa.style.opacity = '0';
  setTimeout(() => {
    qi = (qi + 1) % quotes.length;
    showQuote();
    qt.style.opacity = '1';
    qa.style.opacity = '1';
  }, 400);
}

showQuote();
setInterval(nextQuote, 8000);

function fmt(ms) {
  if (ms <= 0) return '00:00:00';
  const s = Math.floor(ms / 1000);
  const h = Math.floor(s / 3600);
  const m = Math.floor((s % 3600) / 60);
  const sec = s % 60;
  return [h,m,sec].map(n => String(n).padStart(2,'0')).join(':');
}

function tick() {
  const rem = END_MS - Date.now();
  document.getElementById('timer').textContent = fmt(rem);
  if (rem <= 0) {
    document.getElementById('timer').textContent = '00:00:00';
    show('box-done');
    document.getElementById('stop-btn').style.display = 'none';
  }
}

tick();
setInterval(tick, 1000);

let stage = 'idle';

function show(id) {
  ['box-confirm','box-cooldown','box-done'].forEach(b => {
    document.getElementById(b).style.display = b === id ? 'block' : 'none';
  });
}

function phase1() {
  if (stage !== 'idle') return;
  stage = 'confirming';
  document.getElementById('stop-btn').style.display = 'none';
  show('box-confirm');
  setTimeout(() => document.getElementById('confirm-in').focus(), 50);
}

function checkConfirm(val) {
  if (val.trim().toUpperCase() === 'STOP') {
    stage = 'cooldown';
    show('box-cooldown');
    fetch('/api/stop-confirm').then(r => r.json()).then(data => {
      startCooldown(data.cooldown_ms);
    }).catch(() => {
      startCooldown(Date.now() + 120000);
    });
  }
}

function startCooldown(endMs) {
  function update() {
    const rem = endMs - Date.now();
    if (rem <= 0) {
      document.getElementById('cd-display').textContent = 'Done';
      show('box-done');
      stage = 'done';
      return;
    }
    const s = Math.ceil(rem / 1000);
    const m = Math.floor(s / 60);
    const sec = s % 60;
    document.getElementById('cd-display').textContent =
      m + ':' + String(sec).padStart(2,'0');
    setTimeout(update, 500);
  }
  update();
}

function playChime() {
  try {
    const ctx = new (window.AudioContext || window.webkitAudioContext)();
    const notes = [523.25, 659.25, 783.99, 1046.50];
    notes.forEach((freq, i) => {
      const o = ctx.createOscillator();
      const g = ctx.createGain();
      o.connect(g);
      g.connect(ctx.destination);
      o.frequency.value = freq;
      o.type = 'sine';
      const t = ctx.currentTime + i * 0.18;
      g.gain.setValueAtTime(0, t);
      g.gain.linearRampToValueAtTime(0.18, t + 0.05);
      g.gain.exponentialRampToValueAtTime(0.001, t + 0.6);
      o.start(t);
      o.stop(t + 0.65);
    });
  } catch(e) {}
}

const remAtLoad = END_MS - Date.now();
if (remAtLoad > 0) {
  setTimeout(playChime, remAtLoad);
}
</script>
</body>
</html>`
