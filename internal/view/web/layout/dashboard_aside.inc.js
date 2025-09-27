document.addEventListener("DOMContentLoaded", function () {
  const el = document.getElementById("dashboard-aside");
  const key = "dashboard-aside-scroll-position";

  if (!el) return;

  const saveScrollPosition = window.debounce(() => {
    const scrollPosition = el.scrollTop;
    localStorage.setItem(key, scrollPosition);
  }, 200);
  el.addEventListener("scroll", saveScrollPosition);

  const scrollPosition = localStorage.getItem(key);
  if (scrollPosition) {
    el.scrollTop = parseInt(scrollPosition, 10);
  }
});

window.alpineDashboardAsideItem = function (link = "", strict = false) {
  return {
    link,
    strict,
    is_active: false,

    checkActive() {
      if (this.strict) {
        this.is_active = window.location.pathname === this.link;
        return;
      }

      this.is_active = window.location.pathname.startsWith(this.link);
    },

    init() {
      this.checkActive();

      const originalPushState = window.history.pushState;
      window.history.pushState = (...args) => {
        originalPushState.apply(window.history, args);
        this.checkActive();
      };

      const originalReplaceState = window.history.replaceState;
      window.history.replaceState = (...args) => {
        originalReplaceState.apply(window.history, args);
        this.checkActive();
      };
    },
  };
};
