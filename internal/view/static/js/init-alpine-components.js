import { changeThemeButton } from './alpine-components/change-theme-button.js'
import { githubRepoInfo } from './alpine-components/github-repo-info.js'
import { dashboardAsideItem } from './alpine-components/dashboard-aside-item.js'
import { genericSlider } from './alpine-components/generic-slider.js'
import { optionsDropdown } from './alpine-components/options-dropdown.js'

export function initAlpineComponents () {
  document.addEventListener('alpine:init', () => {
    Alpine.data(changeThemeButton.name, changeThemeButton.fn)
    Alpine.data(githubRepoInfo.name, githubRepoInfo.fn)
    Alpine.data(dashboardAsideItem.name, dashboardAsideItem.fn)
    Alpine.data(genericSlider.name, genericSlider.fn)
    Alpine.data(optionsDropdown.name, optionsDropdown.fn)
  })
}
