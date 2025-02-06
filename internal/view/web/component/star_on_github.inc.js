window.alpineStarOnGithub = function () {
  return {
    stars: null,

    async init() {
      const stars = await this.getStars();
      if (stars !== null) {
        this.stars = stars;
      }
    },

    async getStars() {
      const cacheKey = "pbw-gh-stars";

      const cachedJSON = localStorage.getItem(cacheKey);
      if (cachedJSON) {
        const cached = JSON.parse(cachedJSON);
        if (Date.now() - cached.timestamp < 2 * 60 * 1000) {
          return cached.value;
        }
      }

      const url = "https://api.github.com/repos/eduardolat/pgbackweb";
      try {
        const response = await fetch(url);
        if (!response.ok) {
          return null;
        }
        const data = await response.json();
        const value = data.stargazers_count;
        const dataToCache = JSON.stringify({
          value,
          timestamp: Date.now(),
        });
        localStorage.setItem(cacheKey, dataToCache);
        return value;
      } catch {
        return null;
      }
    },
  };
};
