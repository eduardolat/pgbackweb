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

    async init() {
      const stars = await this.getData();
      if (stars !== null) {
        this.stars = stars;
      }
    },

    async getData() {
      const cacheKey = "pbw-support-project-data";

      const cachedJSON = localStorage.getItem(cacheKey);
      if (cachedJSON) {
        const cached = JSON.parse(cachedJSON);

        // Cache for 2 minutes
        if (Date.now() - cached.timestamp < 2 * 60 * 1000) {
          return cached.value;
        }
      }

      const url = "";
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
