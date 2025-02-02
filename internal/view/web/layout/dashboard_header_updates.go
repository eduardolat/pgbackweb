package layout

import (
	"fmt"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func dashboardHeaderUpdates() gomponents.Node {
	return html.A(
		alpine.XData("alpineDashboardHeaderUpdates()"),
		alpine.XCloak(),
		alpine.XShow(fmt.Sprintf(
			"latestRelease !== null && latestRelease !== '%s'",
			config.Version,
		)),

		html.Class("btn btn-warning"),
		html.Href("https://github.com/eduardolat/pgbackweb/releases"),
		html.Target("_blank"),
		lucide.ExternalLink(),
		component.SpanText("Update available"),
		html.Span(
			alpine.XText("'( ' + latestRelease + ' )'"),
		),
	)
}
