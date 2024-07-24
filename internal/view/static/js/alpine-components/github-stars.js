export const githubStars = {
  name: "githubStars",
  fn: () => ({
    stars: "",
    async init() {
      const stars = await this.getGitHubStars()
      if (stars !== null) {
        this.stars = stars
      }
    },
    async getGitHubStars() {
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
    getCachedData(key) {
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
    cacheData(key, value) {
      const data = JSON.stringify({
        value: value,
        timestamp: Date.now(),
      })
      localStorage.setItem(key, data)
    }
  })
}
