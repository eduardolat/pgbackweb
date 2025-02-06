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

    async init() {
      const data = await this.getData();
      if (data !== null) {
        this.data = data;
      }
    },

    async getData() {
      const cacheKey = "pbw-support-project-data";

      const cachedJSON = localStorage.getItem(cacheKey);
      if (cachedJSON) {
        const cached = JSON.parse(cachedJSON);
        // Cache for 2 minutes
        if (Date.now() - cached.timestamp < 2 * 60 * 1000) {
          return cached.data;
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
        return data;
      } catch {
        return null;
      }
    },
  };
};
