export function initDashboardAsideScroll () {
  document.addEventListener('DOMContentLoaded', function () {
    const el = document.getElementById('dashboard-aside')
    const key = 'dashboard-aside-scroll-position'

    if (!el) return

    const saveScrollPosition = window.debounce(
      () => {
        const scrollPosition = el.scrollTop
        localStorage.setItem(key, scrollPosition)
      },
      200
    )
    el.addEventListener('scroll', saveScrollPosition)

    const scrollPosition = localStorage.getItem(key)
    if (scrollPosition) {
      el.scrollTop = parseInt(scrollPosition, 10)
    }
  })
}
