const API_URL = "http://localhost:4862/api/sites";
const SYNC_ALARM_NAME = "focus-sync-rules";
const SYNC_EVERY_MINUTES = 1;
const RULE_ID_OFFSET = 1000;

async function setSyncState(partial) {
  await chrome.storage.local.set({
    ...partial,
    lastSyncAt: Date.now()
  });
}

function escapeRegex(text) {
  return text.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
}

function normalizeDomain(raw) {
  if (typeof raw !== "string") {
    return "";
  }

  let domain = raw.trim().toLowerCase();
  domain = domain.replace(/^https?:\/\//, "");
  domain = domain.split("/")[0];
  domain = domain.split(":")[0];
  domain = domain.replace(/^www\./, "");
  return domain;
}

function buildRule(domain, index) {
  const escapedDomain = escapeRegex(domain);
  return {
    id: RULE_ID_OFFSET + index,
    priority: 1,
    action: {
      type: "redirect",
      redirect: {
        url: `http://localhost:4862/?site=${encodeURIComponent(domain)}`
      }
    },
    condition: {
      regexFilter: `^https?:\\/\\/([^\\/]*\\.)?${escapedDomain}(?::\\d+)?(?:\\/.*)?$`,
      resourceTypes: ["main_frame"]
    }
  };
}

async function fetchActiveSites() {
  try {
    const response = await fetch(API_URL, { cache: "no-store" });
    if (!response.ok) {
      return { online: false, active: false, domains: [] };
    }

    const payload = await response.json();
    if (!payload || !Array.isArray(payload.sites)) {
      return { online: true, active: false, domains: [] };
    }

    const seen = new Set();
    const domains = [];
    for (const site of payload.sites) {
      const domain = normalizeDomain(site);
      if (!domain || seen.has(domain)) {
        continue;
      }
      seen.add(domain);
      domains.push(domain);
    }

    return { online: true, active: payload.active === true, domains };
  } catch {
    return { online: false, active: false, domains: [] };
  }
}

async function replaceDynamicRules(newRules) {
  const existingRules = await chrome.declarativeNetRequest.getDynamicRules();
  const removeRuleIds = existingRules.map((rule) => rule.id);

  await chrome.declarativeNetRequest.updateDynamicRules({
    removeRuleIds,
    addRules: newRules
  });
}

async function syncRules() {
  try {
    const fetched = await fetchActiveSites();
    const domains = fetched.active ? fetched.domains : [];
    const rules = domains.map((domain, index) => buildRule(domain, index));
    await replaceDynamicRules(rules);

    await setSyncState({
      activeDomains: domains,
      syncOk: fetched.online,
      syncMessage: !fetched.online
        ? "Focus app is offline."
        : domains.length > 0
          ? "Redirect rules are active."
          : "No active blocked domains.",
      focusOnline: fetched.online
    });

    return { ok: true, domains };
  } catch (error) {
    const message = error && error.message ? error.message : "Failed to update rules.";
    await setSyncState({
      activeDomains: [],
      syncOk: false,
      syncMessage: message,
      focusOnline: false
    });
    return { ok: false, error: message };
  }
}

async function setupAlarm() {
  await chrome.alarms.clear(SYNC_ALARM_NAME);
  await chrome.alarms.create(SYNC_ALARM_NAME, {
    delayInMinutes: 1,
    periodInMinutes: SYNC_EVERY_MINUTES
  });
}

chrome.runtime.onInstalled.addListener(async () => {
  await setupAlarm();
  await syncRules();
});

chrome.runtime.onStartup.addListener(async () => {
  await setupAlarm();
  await syncRules();
});

chrome.alarms.onAlarm.addListener(async (alarm) => {
  if (alarm.name !== SYNC_ALARM_NAME) {
    return;
  }
  await syncRules();
});

chrome.runtime.onMessage.addListener((message, _sender, sendResponse) => {
  if (!message || typeof message.type !== "string") {
    return;
  }

  if (message.type === "focus.syncNow") {
    syncRules()
      .then((result) => sendResponse(result))
      .catch((error) => {
        sendResponse({ ok: false, error: error && error.message ? error.message : "Unknown error" });
      });
    return true;
  }

  if (message.type === "focus.getStatus") {
    chrome.storage.local.get(["lastSyncAt", "activeDomains", "syncOk", "syncMessage", "focusOnline"], (data) => {
      sendResponse({
        lastSyncAt: data.lastSyncAt || 0,
        activeDomains: Array.isArray(data.activeDomains) ? data.activeDomains : [],
        syncOk: data.syncOk === true,
        syncMessage: data.syncMessage || "Waiting for first sync.",
        focusOnline: data.focusOnline === true
      });
    });
    return true;
  }
});
