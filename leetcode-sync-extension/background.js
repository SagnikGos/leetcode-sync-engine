console.log("LeetCode Sync: background service worker loaded");

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
    if (message.type === "LEETCODE_SYNC_SUBMISSION") {
        console.log("LeetCode Sync: Submission received in background");

        console.log("Payload:", message.payload);

        // We will send to backend in next chunk

        sendResponse({ status: "received" });
    }

    // Return true to indicate asynchronous response if needed mostly used with fetch
    return true;
});
