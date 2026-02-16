console.log("LeetCode Sync: background service worker loaded");

const BACKEND_URL = "http://localhost:8080/api/submission";

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
    if (message.type === "LEETCODE_SYNC_SUBMISSION") {
        console.log("LeetCode Sync: Submission received in background");

        sendToBackend(message.payload)
            .then(response => {
                console.log("LeetCode Sync: Backend success:", response);
                sendResponse({ status: "success" });
            })
            .catch(error => {
                console.error("LeetCode Sync: Backend error:", error);
                sendResponse({ status: "error" });
            });

        return true; // keep channel open for async response
    }
});

async function sendToBackend(payload) {
    try {
        const response = await fetch(BACKEND_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(payload)
        });

        if (!response.ok) {
            throw new Error(`Backend returned ${response.status}`);
        }

        return await response.json();
    } catch (error) {
        throw error;
    }
}
