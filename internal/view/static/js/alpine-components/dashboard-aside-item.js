export const dashboardAsideItem = {
  name: 'dashboardAsideItem',
  fn: (link = '', strict = false) => ({
    link,
    strict,
    is_active: false,
    init () {
      if (this.strict) {
        this.is_active = window.location.pathname === this.link
        return
      }

      this.is_active = window.location.pathname.startsWith(this.link)
    }
  })
}
