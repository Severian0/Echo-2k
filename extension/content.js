// content.js

console.log("[Echo-2k] content script injected on:", location.href);

// A simple state object you can extend as you discover more data
const echo2kState = {
    pageUrl: location.href,
    pageTitle: document.title,
    userEmail: null,
    // add more fields as you figure out what you need
};

// Example function to grab some info from the DOM
function initEcho2k() {
    console.log("[Echo-2k] initEcho2k() running...");

    const el = document.querySelector(selector);

    if (el.tagName.toLowerCase() === "a" && el.href.startsWith("mailto:")) {
        echo2kState.userEmail = el.href.replace("mailto:", "").trim();
    } else {
        echo2kState.userEmail = el.textContent.trim();
    }

  console.log("[Echo-2k] current state:", echo2kState);
}

// Because run_at is document_idle, DOM should already be there.
// But to be extra explicit, you can still attach to DOMContentLoaded if you want:

if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", initEcho2k);
} else {
  initEcho2k();
}

// You can now use `echo2kState` from this content script file going forward.
// Any other functions in content.js can read/write it.
