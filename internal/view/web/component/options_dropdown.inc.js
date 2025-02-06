window.alpineOptionsDropdown = function () {
  return {
    isOpen: false,
    buttonEl: null,
    contentEl: null,
    closeTimeout: null,

    init() {
      this.buttonEl = this.$refs.button;
      this.contentEl = this.$refs.content;
    },

    open() {
      this.isOpen = true;
      this.contentEl.classList.remove("hidden");
      this.positionContent();

      if (this.closeTimeout) {
        clearTimeout(this.closeTimeout);
        this.closeTimeout = null;
      }
    },

    close() {
      this.closeTimeout = setTimeout(() => {
        this.isOpen = false;
        this.contentEl.classList.add("hidden");
      }, 200);
    },

    positionContent() {
      const buttonRect = this.buttonEl.getBoundingClientRect();
      const contentHeight = this.contentEl.offsetHeight;
      const windowHeight = window.innerHeight;
      const moreSpaceBelow =
        (windowHeight - buttonRect.bottom) > buttonRect.top;

      this.contentEl.style.left = `${buttonRect.left}px`;

      if (moreSpaceBelow) {
        this.contentEl.style.top = `${buttonRect.bottom}px`;
      } else {
        this.contentEl.style.top = `${buttonRect.top - contentHeight}px`;
      }
    },
  };
};
