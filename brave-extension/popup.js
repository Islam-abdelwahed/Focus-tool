function formatTime(ts) {
  if (!ts) {
    return "--";
  }

  const date = new Date(ts);
  const now = Date.now();
  const diffSec = Math.max(0, Math.floor((now - ts) / 1000));

  if (diffSec < 5) {
    return "now";
  }
  if (diffSec < 60) {
    return `${diffSec}s`;
  }
  if (diffSec < 3600) {
    return `${Math.floor(diffSec / 60)}m`;
  }
  return date.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
}

function updateBadge(online) {
  const badge = document.getElementById("api-badge");
  badge.classList.remove("online", "offline");

  if (online) {
    badge.classList.add("online");
    badge.textContent = "Focus Online";
  } else {
    badge.classList.add("offline");
    badge.textContent = "Focus Offline";
  }
}

function renderDomains(domains) {
  const list = document.getElementById("domains-list");
  const empty = document.getElementById("domains-empty");

  list.innerHTML = "";
  if (!domains.length) {
    list.hidden = true;
    empty.hidden = false;
    return;
  }

  domains.forEach((domain) => {
    const item = document.createElement("li");
    item.textContent = domain;
    list.appendChild(item);
  });

  empty.hidden = true;
  list.hidden = false;
}

async function getStatus() {
  return new Promise((resolve) => {
    chrome.runtime.sendMessage({ type: "focus.getStatus" }, (response) => {
      if (chrome.runtime.lastError) {
        resolve({
          activeDomains: [],
          lastSyncAt: 0,
          syncOk: false,
          syncMessage: chrome.runtime.lastError.message || "Unable to read status.",
          focusOnline: false
        });
        return;
      }
      resolve(response || {
        activeDomains: [],
        lastSyncAt: 0,
        syncOk: false,
        syncMessage: "No status available.",
        focusOnline: false
      });
    });
  });
}

async function syncNow() {
  return new Promise((resolve) => {
    chrome.runtime.sendMessage({ type: "focus.syncNow" }, (response) => {
      if (chrome.runtime.lastError) {
        resolve({ ok: false, error: chrome.runtime.lastError.message || "Failed to sync." });
        return;
      }
      resolve(response || { ok: false, error: "No response from background." });
    });
  });
}

async function refreshView() {
  const status = await getStatus();
  document.getElementById("status-msg").textContent = status.syncMessage || "No status available.";
  document.getElementById("domains-count").textContent = String((status.activeDomains || []).length);
  document.getElementById("last-sync").textContent = formatTime(status.lastSyncAt || 0);
  updateBadge(status.focusOnline === true);
  renderDomains(Array.isArray(status.activeDomains) ? status.activeDomains : []);
}

async function handleSync() {
  const button = document.getElementById("sync-btn");
  button.disabled = true;
  button.textContent = "Syncing...";

  await syncNow();
  await refreshView();

  button.disabled = false;
  button.textContent = "Sync Now";
}

document.getElementById("sync-btn").addEventListener("click", handleSync);
document.getElementById("refresh-btn").addEventListener("click", refreshView);

refreshView();
setInterval(refreshView, 15000);
