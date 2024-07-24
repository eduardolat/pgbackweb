export const changeThemeButton = {
  name: "changeThemeButton",
  fn: () => ({
    theme: "",

    getCurrentTheme() {
      const el = document.querySelector("html")
      const theme = el.getAttribute("data-theme")
      if (theme) {
        this.theme = theme
        return
      }
      this.theme = "system"
    },

    init() {
      setTimeout(() => {
        this.getCurrentTheme()
      }, 200)
    }
  })
}
