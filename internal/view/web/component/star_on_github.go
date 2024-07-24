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
		alpine.XData("githubStars"),
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
