export function initDialogs() {
  /**
   * Shows an alert dialog
   * @param {string} text - The text to display
   * @returns {Promise<{isConfirmed: boolean, isDismissed: boolean}>}
   */
  async function customAlert(text) {
    return showDialog(text, false);
  }

  /**
   * Shows a confirmation dialog
   * @param {string} text - The text to display
   * @returns {Promise<{isConfirmed: boolean, isDismissed: boolean}>}
   */
  async function customConfirm(text) {
    return showDialog(text, true);
  }

  /**
   * Shows a dialog
   * @param {string} text - The text to display
   * @param {boolean} isConfirm - True for confirm dialog, false for alert
   */
  function showDialog(text, isConfirm) {
    return new Promise((resolve) => {
      const dialogId = "dialog-" + Date.now();
      const container = createDialog(dialogId, text, isConfirm, resolve);
      document.body.appendChild(container);

      // Fade in
      requestAnimationFrame(() => {
        container.style.opacity = "0";
        container.classList.remove("hidden");
        requestAnimationFrame(() => {
          container.style.transition = "opacity 0.15s ease-in-out";
          container.style.opacity = "1";
        });
      });
    });
  }

  /**
   * Creates the dialog HTML
   */
  function createDialog(dialogId, text, isConfirm, resolve) {
    // Container
    const container = document.createElement("div");
    container.id = dialogId;
    container.className =
      "hidden !p-0 !m-0 w-[100dvw] h-[100dvh] fixed left-0 top-0 z-[1000]";

    // Backdrop
    const backdrop = document.createElement("div");
    backdrop.className = "bg-black opacity-25 !w-full !h-full z-[1001]";
    backdrop.onclick = () => closeDialog(dialogId, resolve, !isConfirm);

    // Dialog box
    const dialogBox = document.createElement("div");
    dialogBox.className =
      "absolute z-[1002] top-[50%] left-[50%] translate-y-[-50%] translate-x-[-50%] " +
      "max-w-[calc(100dvw-30px)] bg-base-100 rounded-box p-6 w-[400px] shadow-xl";

    // Icon
    const iconPath = isConfirm
      ? "M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
      : "M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z";
    const iconColor = isConfirm ? "text-warning" : "text-info";

    dialogBox.innerHTML = `
      <div class="flex justify-center mb-4">
        <svg xmlns="http://www.w3.org/2000/svg" class="size-[100px] ${iconColor}" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="${iconPath}" />
        </svg>
      </div>
      <div class="text-center text-lg mb-6 break-words">${text}</div>
      <div class="flex justify-center gap-3">
        ${
          isConfirm
            ? '<button class="btn btn-error" data-action="cancel">Cancel</button><button class="btn btn-primary" data-action="confirm">Confirm</button>'
            : '<button class="btn btn-primary" data-action="ok">Okay</button>'
        }
      </div>
    `;

    // Button event listeners
    const buttons = dialogBox.querySelectorAll("button");
    buttons.forEach((btn) => {
      btn.onclick = () => {
        const action = btn.getAttribute("data-action");
        const confirmed = action === "confirm" || action === "ok";
        closeDialog(dialogId, resolve, confirmed);
      };
    });

    // Autofocus
    setTimeout(() => {
      const focusSelector = isConfirm
        ? '[data-action="cancel"]'
        : '[data-action="ok"]';
      dialogBox.querySelector(focusSelector)?.focus();
    }, 150);

    container.appendChild(backdrop);
    container.appendChild(dialogBox);

    // ESC key handler
    const handleEsc = (e) => {
      if (e.key === "Escape") {
        closeDialog(dialogId, resolve, !isConfirm);
        document.removeEventListener("keydown", handleEsc);
      }
    };
    document.addEventListener("keydown", handleEsc);

    return container;
  }

  /**
   * Closes the dialog with fade out
   */
  function closeDialog(dialogId, resolve, confirmed) {
    const dialog = document.getElementById(dialogId);
    if (dialog) {
      dialog.style.opacity = "0";
      setTimeout(() => {
        dialog.remove();
        resolve({ isConfirmed: confirmed, isDismissed: !confirmed });
      }, 150);
    } else {
      resolve({ isConfirmed: confirmed, isDismissed: !confirmed });
    }
  }

  window.customAlert = customAlert;
  window.customConfirm = customConfirm;
}
