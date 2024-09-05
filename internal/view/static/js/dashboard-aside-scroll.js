export function initDashboardAsideScroll () {
  const el = document.getElementById('dashboard-aside')
  const key = 'dashboard-aside-scroll-position'

  if (!el) return

  const saveScrollPosition = window.debounce(
    () => {
      const scrollPosition = el.scrollTop
      localStorage.setItem(key, scrollPosition)
      console.log(scrollPosition)
    },
    200
  )
  el.addEventListener('scroll', saveScrollPosition)

  document.addEventListener('DOMContentLoaded', function () {
    const scrollPosition = localStorage.getItem(key)
    if (scrollPosition) {
      el.scrollTop = parseInt(scrollPosition, 10)
    }
  })
}
