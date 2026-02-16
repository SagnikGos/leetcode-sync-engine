const script = document.createElement("script");
script.src = chrome.runtime.getURL("inject.js");
script.onload = function () {
    this.remove();
};
(document.head || document.documentElement).appendChild(script);

console.log("LeetCode Sync: content script loaded (observer mode)");

let submissionProcessed = false;

function getSlug() {
    const parts = window.location.pathname.split("/");
    return parts[2] || null;
}

function getTitle() {
    // Try primary selector
    let el = document.querySelector('[data-cy="question-title"]');
    if (el) return el.innerText.trim();

    // Fallback 1
    el = document.querySelector('.text-title-large');
    if (el) return el.innerText.trim();

    // Fallback 2 — first h1
    const h1 = document.querySelector('h1');
    if (h1) return h1.innerText.trim();

    return null;
}

function getDifficulty() {
    const difficultyEls = document.querySelectorAll("div");

    for (const el of difficultyEls) {
        const text = el.innerText?.trim();
        if (text === "Easy" || text === "Medium" || text === "Hard") {
            return text;
        }
    }
    return null;
}

function getDescriptionHTML() {
    // Strict selector for new UI (confirmed stable)
    const desc = document.querySelector(
        '[data-track-load="description_content"]'
    );

    if (!desc) {
        //        console.log("LeetCode Sync: Description container not found");
        return null;
    }

    return desc.innerHTML.trim();
}

function getLanguage() {
    // Try to find visible language dropdown button text by checking all buttons
    // This avoids fragile ID selectors and dynamic attributes
    const buttons = document.querySelectorAll("button");

    for (const btn of buttons) {
        const text = btn.innerText?.trim();

        if (!text) continue;

        // Check common LeetCode languages
        if (
            text === "C++" ||
            text === "Java" ||
            text === "Python" ||
            text === "Python3" ||
            text === "JavaScript" ||
            text === "TypeScript" ||
            text === "Go" ||
            text === "C#" ||
            text === "Rust" ||
            text === "Ruby" ||
            text === "Swift" ||
            text === "Kotlin" ||
            text === "PHP" ||
            text === "Dart" ||
            text === "Scala" ||
            text === "Racket" ||
            text === "Erlang" ||
            text === "Elixir"
        ) {
            return text;
        }
    }

    return "unknown";
}

function extractMetadata() {
    const metadata = {
        slug: getSlug(),
        title: getTitle(),
        difficulty: getDifficulty(),
        description_html: getDescriptionHTML(), // Log full HTML for validation
        language: getLanguage(),
        timestamp: new Date().toISOString()
    };

    return metadata;
}

// Listen for code response from inject.js
window.addEventListener("message", function (event) {
    if (event.source !== window) return;

    if (event.data.type === "LEETCODE_SYNC_CODE_RESPONSE") {
        const code = event.data.code;

        if (!code) {
            console.log("LeetCode Sync: No code found");
            return;
        }

        const metadata = extractMetadata();

        const submissionPayload = {
            ...metadata,
            code: code
        };

        // Send to background script
        chrome.runtime.sendMessage({
            type: "LEETCODE_SYNC_SUBMISSION",
            payload: submissionPayload
        });
    }
});

function checkForAccepted() {
    if (submissionProcessed) return;

    // Correct selector for result container panel inside virtualized DOM
    const resultContainer = document.querySelector(
        '[data-e2e-locator="submission-result"]'
    );

    if (!resultContainer) return;

    const resultText = resultContainer.innerText;

    if (resultText.includes("Accepted")) {
        console.log("LeetCode Sync: Accepted detected");

        submissionProcessed = true;

        // Request code from page context with a small delay to ensure Monaco is ready
        setTimeout(() => {
            window.postMessage({
                type: "LEETCODE_SYNC_GET_CODE"
            }, "*");
        }, 500);
    }
}

const observer = new MutationObserver(() => {
    checkForAccepted();
});

observer.observe(document.body, {
    childList: true,
    subtree: true
});
