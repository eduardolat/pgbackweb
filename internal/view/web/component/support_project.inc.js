// EXAMPLE OF DATA
// {
//   "referralLinks": [
//     {
//       "name": "...",
//       "logo": "...",
//       "link": "...",
//       "description": "..."
//     },
//   ],
//   "sponsors": {
//     "link": "...",
//     "bronze": [
//       {
//         "name": "...",
//         "logo": "...",
//         "link": "..."
//       }
//     ]
//   }
// }

window.alpineSupportProjectData = function () {
  return {
    data: null,

    get isLoading() {
      return this.data === null;
    },

    get sponsorsLink() {
      return this.data?.sponsors?.link ?? "";
    },

    get referralLinks() {
      return this.data?.referralLinks ?? [];
    },

    get bronzeSponsors() {
      return this.data?.sponsors?.bronze ?? [];
    },

    get silverSponsors() {
      return this.data?.sponsors?.silver ?? [];
    },

    get goldSponsors() {
      return this.data?.sponsors?.gold ?? [];
    },

    async init() {
      const data = await this.getData();
      if (data !== null) {
        this.data = data;
      }
    },

    prefixImagePath(path) {
      // If the path starts with / and doesn't start with http, add the path prefix
      if (path && path.startsWith("/") && !path.startsWith("http")) {
        return window.PBW_PATH_PREFIX + path;
      }
      return path;
    },

    processData(data) {
      // Add path prefix to all image URLs
      if (data.referralLinks) {
        data.referralLinks = data.referralLinks.map((link) => ({
          ...link,
          logo: this.prefixImagePath(link.logo),
        }));
      }

      if (data.sponsors) {
        if (data.sponsors.gold) {
          data.sponsors.gold = data.sponsors.gold.map((sponsor) => ({
            ...sponsor,
            logo: this.prefixImagePath(sponsor.logo),
          }));
        }
        if (data.sponsors.silver) {
          data.sponsors.silver = data.sponsors.silver.map((sponsor) => ({
            ...sponsor,
            logo: this.prefixImagePath(sponsor.logo),
          }));
        }
        if (data.sponsors.bronze) {
          data.sponsors.bronze = data.sponsors.bronze.map((sponsor) => ({
            ...sponsor,
            logo: this.prefixImagePath(sponsor.logo),
          }));
        }
      }

      return data;
    },

    async getData() {
      const cacheKey = "pbw-support-project-data";

      const cachedJSON = localStorage.getItem(cacheKey);
      if (cachedJSON) {
        const cached = JSON.parse(cachedJSON);
        // Cache for 2 minutes
        if (Date.now() - cached.timestamp < 2 * 60 * 1000) {
          return this.processData(cached.data);
        }
      }

      const url =
        "https://raw.githubusercontent.com/eduardolat/pgbackweb/refs/heads/develop/assets/support-project-v1.json";
      try {
        const response = await fetch(url);
        if (!response.ok) {
          return null;
        }
        const data = await response.json();
        const dataToCache = JSON.stringify({
          data,
          timestamp: Date.now(),
        });
        localStorage.setItem(cacheKey, dataToCache);
        return this.processData(data);
      } catch {
        return null;
      }
    },
  };
};
