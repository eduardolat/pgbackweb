package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func StarOnGithub(size size) gomponents.Node {
	return html.A(
		alpine.XData(`{
			stars: "",
			async init() {
				const stars = await this.getGitHubStars();
				if (stars !== null) {
					this.stars = stars;
				}
			},
			async getGitHubStars() {
				const cacheKey = 'pbw_gh_stars';
				const cachedData = this.getCachedData(cacheKey);
				if (cachedData !== null) {
					return cachedData;
				}

				const url = 'https://api.github.com/repos/eduardolat/pgbackweb';
				try {
					const response = await fetch(url);
					if (!response.ok) {
						return null;
					}
					const data = await response.json();
					this.cacheData(cacheKey, data.stargazers_count);
					return data.stargazers_count;
				} catch {
					return null;
				}
			},
			getCachedData(key) {
				const cachedJSON = localStorage.getItem(key);
				if (!cachedJSON) {
					return null;
				}
				const cached = JSON.parse(cachedJSON);
				if (Date.now() - cached.timestamp < 2 * 60 * 1000) {
					return cached.value;
				}
				return null;
			},
			cacheData(key, value) {
				const data = JSON.stringify({
					value: value,
					timestamp: Date.now(),
				});
				localStorage.setItem(key, data);
			}
		}`),
		alpine.XCloak(),
		components.Classes{
			"btn btn-neutral": true,
			"btn-sm":          size == SizeSm,
			"btn-lg":          size == SizeLg,
		},
		html.Href("https://github.com/eduardolat/pgbackweb"),
		html.Target("_blank"),
		lucide.Github(),
		SpanText("Star on Github"),
		html.Span(
			alpine.XShow("stars"),
			alpine.XText("'( ' + stars + ' )'"),
		),
	)
}
