window.alpineDashboardHeaderUpdates = function () {
  return {
    latestRelease: null,

    async init () {
      const latestRelease = await this.getLatestRelease()
      if (latestRelease !== null) {
        this.latestRelease = latestRelease
      }
    },

    async getLatestRelease () {
      const cacheKey = 'pbw-gh-last-release'

      const cachedJSON = localStorage.getItem(cacheKey)
      if (cachedJSON) {
        const cached = JSON.parse(cachedJSON)
        if (Date.now() - cached.timestamp < 2 * 60 * 1000) {
          return cached.value
        }
      }

      const url = 'https://api.github.com/repos/eduardolat/pgbackweb/releases/latest'
      try {
        const response = await fetch(url)
        if (!response.ok) {
          return null
        }
        const data = await response.json()
        const value = data.name
        const dataToCache = JSON.stringify({
          value,
          timestamp: Date.now()
        })
        localStorage.setItem(cacheKey, dataToCache)
        return value
      } catch {
        return null
      }
    }
  }
}
