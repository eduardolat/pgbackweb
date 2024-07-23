var toastQueue = [];

const toastCfg = {
  duration: 5000,
  ripple: true,
  position: { x: 'right', y: 'bottom' },
  dismissible: true,
  ripple: false
};

const infiniteToastCfg = {
  ...toastCfg,
  duration: 0
};

const toaster = {
  success: (message) => {
    toastQueue.push({ type: "success", message, config: toastCfg });
  },
  error: (message) => {
    toastQueue.push({ type: "error", message, config: toastCfg });
  },
  successInfinite: (message) => {
    toastQueue.push({ type: "success", message, config: infiniteToastCfg });
  },
  errorInfinite: (message) => {
    toastQueue.push({ type: "error", message, config: infiniteToastCfg });
  }
};

document.addEventListener('DOMContentLoaded', function () {
  var notyf = new Notyf();

  toastQueue.forEach(item => {
    notyf.open({ type: item.type, message: item.message, ...item.config });
  });

  toastQueue = [];

  toaster.success = (message) => {
    notyf.open({ type: "success", message, ...toastCfg });
  };
  toaster.error = (message) => {
    notyf.open({ type: "error", message, ...toastCfg });
  };
  toaster.successInfinite = (message) => {
    notyf.open({ type: "success", message, ...infiniteToastCfg });
  };
  toaster.errorInfinite = (message) => {
    notyf.open({ type: "error", message, ...infiniteToastCfg });
  };
});
