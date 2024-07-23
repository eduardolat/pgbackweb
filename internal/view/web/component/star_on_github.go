package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func StarOnGithub(size size) gomponents.Node {
	return html.A(
		components.Classes{
			"btn btn-neutral": true,
			"btn-sm":          size == SizeSm,
			"btn-lg":          size == SizeLg,
		},
		html.Href("https://github.com/eduardolat/pgbackweb"),
		html.Target("_blank"),
		lucide.Github(),
		SpanText("Star on GitHub"),
	)
}
