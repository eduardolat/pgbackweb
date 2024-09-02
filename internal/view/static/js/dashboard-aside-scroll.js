export function initDashboardAsideScroll () {
  const el = document.getElementById('dashboard-aside')
  const key = 'dashboard-aside-scroll-position'

  if (!el) return

  window.addEventListener('beforeunload', function () {
    const scrollPosition = el.scrollTop
    localStorage.setItem(key, scrollPosition)
  })

  document.addEventListener('DOMContentLoaded', function () {
    const scrollPosition = localStorage.getItem(key)
    if (scrollPosition) {
      el.scrollTop = parseInt(scrollPosition, 10)
    }
  })
}
