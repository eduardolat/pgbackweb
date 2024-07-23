export function initHTMXTriggers() {
  const triggers = {
    ctm_alert: function (evt) {
      const message = decodeURIComponent(evt.detail.value);
      alert(message);
    },
    ctm_alert_with_refresh: function (evt) {
      const message = decodeURIComponent(evt.detail.value);
      alert(message);
      location.reload();
    },
    ctm_alert_with_redirect: function (evt) {
      const payload = decodeURIComponent(evt.detail.value);
      const parts = payload.split('-::-::-');
      if (parts.length !== 2) {
        return;
      }
      const message = parts[0];
      const url = parts[1];

      alert(message);
      location.href = url;
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
  }

  for (const key in triggers) {
    document.addEventListener(key, triggers[key])
  }
}
