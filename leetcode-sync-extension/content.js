console.log("LeetCode Sync: content script loaded (observer mode)");

let submissionProcessed = false;

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

        alert("Accepted submission detected!");
    }
}

const observer = new MutationObserver(() => {
    checkForAccepted();
});

observer.observe(document.body, {
    childList: true,
    subtree: true
});
