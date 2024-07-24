import { changeThemeButton } from './alpine-components/change-theme-button.js'
import { githubStars } from './alpine-components/github-stars.js'
import { dashboardAsideItem } from './alpine-components/dashboard-aside-item.js'

export function initAlpineComponents () {
  document.addEventListener('alpine:init', () => {
    Alpine.data(changeThemeButton.name, changeThemeButton.fn)
    Alpine.data(githubStars.name, githubStars.fn)
    Alpine.data(dashboardAsideItem.name, dashboardAsideItem.fn)
  })
}
