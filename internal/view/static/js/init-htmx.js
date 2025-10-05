export function initHTMX() {
  const triggers = {
    ctm_alert: function (evt) {
      const message = decodeURIComponent(evt.detail.value);
      window.swalAlert(message);
    },
    ctm_alert_with_refresh: function (evt) {
      const message = decodeURIComponent(evt.detail.value);
      window.swalAlert(message).then(() => {
        location.reload();
      });
    },
    ctm_alert_with_redirect: function (evt) {
      const payload = decodeURIComponent(evt.detail.value);
      const parts = payload.split("-::-::-");
      if (parts.length !== 2) {
        return;
      }
      const message = parts[0];
      const url = parts[1];

      window.swalAlert(message).then(() => {
        location.href = url;
      });
    },
    ctm_toast_success: function (evt) {
      const message = decodeURIComponent(evt.detail.value);
      toaster.success(message);
    },
    ctm_toast_error: function (evt) {
      const message = decodeURIComponent(evt.detail.value);
      toaster.error(message);
    },
    ctm_toast_success_infinite: function (evt) {
      const message = decodeURIComponent(evt.detail.value);
      toaster.successInfinite(message);
    },
    ctm_toast_error_infinite: function (evt) {
      const message = decodeURIComponent(evt.detail.value);
      toaster.errorInfinite(message);
    },
  };

  for (const key in triggers) {
    document.addEventListener(key, triggers[key]);
  }

  // Add trigger to use custom dialogs for confirms
  document.addEventListener("htmx:confirm", function (e) {
    if (!e.detail.target.hasAttribute("hx-confirm")) return;

    e.preventDefault();
    window.swalConfirm(e.detail.question).then(function (result) {
      if (result.isConfirmed) e.detail.issueRequest(true);
    });
  });

  // This fixes this issue:
  // https://stackoverflow.com/questions/73658449/htmx-request-not-firing-when-hx-attributes-are-added-dynamically-from-javascrip
  document.addEventListener("DOMContentLoaded", function () {
    const observer = new MutationObserver((mutations) => {
      mutations.forEach((mutation) => {
        mutation.addedNodes.forEach((node) => {
          if (node.nodeType === 1 && !node["htmx-internal-data"]) {
            htmx.process(node);
          }
        });
      });
    });
    observer.observe(document, { childList: true, subtree: true });
  });
}
