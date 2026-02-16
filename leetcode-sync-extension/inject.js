(function () {
    window.addEventListener("message", function (event) {
        if (event.source !== window) return;

        if (event.data.type === "LEETCODE_SYNC_GET_CODE") {
            let code = null;

            try {
                if (window.monaco && window.monaco.editor) {
                    const models = window.monaco.editor.getModels();
                    if (models.length > 0) {
                        code = models[0].getValue();
                    }
                }
            } catch (err) {
                console.error("LeetCode Sync: Monaco access error", err);
            }

            window.postMessage({
                type: "LEETCODE_SYNC_CODE_RESPONSE",
                code: code
            }, "*");
        }
    });
})();
