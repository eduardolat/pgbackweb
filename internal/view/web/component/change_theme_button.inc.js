window.alpineChangeThemeButton = function () {
  return {
    theme: "",

    loadTheme() {
      const theme = window.getTheme();
      this.theme = theme || "system";
    },

    setTheme(theme) {
      window.setTheme(theme);
      this.theme = theme || "system";
    },

    init() {
      this.loadTheme();
    },
  };
};
