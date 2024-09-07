export const changeThemeButton = {
  name: 'changeThemeButton',
  fn: () => ({
    theme: '',

    loadTheme () {
      const theme = window.getTheme()
      this.theme = theme || 'system'
    },

    setTheme (theme) {
      window.setTheme(theme)
      this.theme = theme || 'system'
    },

    init () {
      this.loadTheme()
    }
  })
}
