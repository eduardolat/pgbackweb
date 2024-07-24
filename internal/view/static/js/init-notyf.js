export function initNotyf () {
  let toastQueue = []

  const toastCfg = {
    duration: 5000,
    ripple: false,
    position: { x: 'right', y: 'bottom' },
    dismissible: true
  }

  const infiniteToastCfg = {
    ...toastCfg,
    duration: 0
  }

  window.toaster = {
    success: (message) => {
      toastQueue.push({ type: 'success', message, config: toastCfg })
    },
    error: (message) => {
      toastQueue.push({ type: 'error', message, config: toastCfg })
    },
    successInfinite: (message) => {
      toastQueue.push({ type: 'success', message, config: infiniteToastCfg })
    },
    errorInfinite: (message) => {
      toastQueue.push({ type: 'error', message, config: infiniteToastCfg })
    }
  }

  document.addEventListener('DOMContentLoaded', function () {
    const notyf = new Notyf()

    toastQueue.forEach(item => {
      notyf.open({ type: item.type, message: item.message, ...item.config })
    })

    toastQueue = []

    window.toaster.success = (message) => {
      notyf.open({ type: 'success', message, ...toastCfg })
    }
    window.toaster.error = (message) => {
      notyf.open({ type: 'error', message, ...toastCfg })
    }
    window.toaster.successInfinite = (message) => {
      notyf.open({ type: 'success', message, ...infiniteToastCfg })
    }
    window.toaster.errorInfinite = (message) => {
      notyf.open({ type: 'error', message, ...infiniteToastCfg })
    }
  })
}
