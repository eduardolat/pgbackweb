export function initSweetAlert2() {
  // Docs at https://sweetalert2.github.io/#configuration
  const defaultConfig = {
    icon: "info",
    confirmButtonText: "Okay",
    cancelButtonText: "Cancel",
    customClass: {
      popup: "rounded-box bg-base-100 text-base-content",
      confirmButton: "btn btn-primary",
      denyButton: "btn btn-warning",
      cancelButton: "btn btn-error",
    },
  };

  async function swalAlert(text) {
    return await Swal.fire({
      ...defaultConfig,
      title: text,
    });
  }

  async function swalConfirm(text) {
    return await Swal.fire({
      ...defaultConfig,
      icon: "question",
      title: text,
      confirmButtonText: "Confirm",
      showCancelButton: true,
    });
  }

  window.swalAlert = swalAlert;
  window.swalConfirm = swalConfirm;
}
