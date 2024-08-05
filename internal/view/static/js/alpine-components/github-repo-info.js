export const githubRepoInfo = {
  name: 'githubRepoInfo',
  fn: () => ({
    stars: '',
    latestRelease: '',

    async init () {
      const stars = await this.getStars()
      if (stars !== null) {
        this.stars = stars
      }

      const latestRelease = await this.getLatestRelease()
      if (latestRelease !== null) {
        this.latestRelease = latestRelease
      }
    },
    async getStars () {
      const cacheKey = 'pbw_gh_stars'
      const cachedData = this.getCachedData(cacheKey)
      if (cachedData !== null) {
        return cachedData
      }

      const url = 'https://api.github.com/repos/eduardolat/pgbackweb'
      try {
        const response = await fetch(url)
        if (!response.ok) {
          return null
        }
        const data = await response.json()
        this.cacheData(cacheKey, data.stargazers_count)
        return data.stargazers_count
      } catch {
        return null
      }
    },
    async getLatestRelease () {
      const cacheKey = 'pbw_gh_last_release'
      const cachedData = this.getCachedData(cacheKey)
      if (cachedData !== null) {
        return cachedData
      }

      const url = 'https://api.github.com/repos/eduardolat/pgbackweb/releases/latest'
      try {
        const response = await fetch(url)
        if (!response.ok) {
          return null
        }
        const data = await response.json()
        this.cacheData(cacheKey, data.name)
        return data.name
      } catch {
        return null
      }
    },
    getCachedData (key) {
      const cachedJSON = localStorage.getItem(key)
      if (!cachedJSON) {
        return null
      }
      const cached = JSON.parse(cachedJSON)
      if (Date.now() - cached.timestamp < 2 * 60 * 1000) {
        return cached.value
      }
      return null
    },
    cacheData (key, value) {
      const data = JSON.stringify({
        value,
        timestamp: Date.now()
      })
      localStorage.setItem(key, data)
    }
  })
}
