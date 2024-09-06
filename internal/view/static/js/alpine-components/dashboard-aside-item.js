export const dashboardAsideItem = {
  name: 'dashboardAsideItem',
  fn: (link = '', strict = false) => ({
    link,
    strict,
    is_active: false,

    checkActive () {
      if (this.strict) {
        this.is_active = window.location.pathname === this.link
        return
      }

      this.is_active = window.location.pathname.startsWith(this.link)
    },

    init () {
      this.checkActive()

      const originalPushState = window.history.pushState
      window.history.pushState = (...args) => {
        originalPushState.apply(window.history, args)
        this.checkActive()
      }

      const originalReplaceState = window.history.replaceState
      window.history.replaceState = (...args) => {
        originalReplaceState.apply(window.history, args)
        this.checkActive()
      }
    }
  })
}
