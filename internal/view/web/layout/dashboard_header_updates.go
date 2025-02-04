package layout

import (
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func dashboardHeaderUpdates() nodx.Node {
	return nodx.A(
		alpine.XData("alpineDashboardHeaderUpdates()"),
		alpine.XCloak(),
		alpine.XShow(fmt.Sprintf(
			"latestRelease !== null && latestRelease !== '%s'",
			config.Version,
		)),

		nodx.Class("btn btn-warning"),
		nodx.Href("https://github.com/eduardolat/pgbackweb/releases"),
		nodx.Target("_blank"),
		lucide.ExternalLink(),
		component.SpanText("Update available"),
		nodx.SpanEl(
			alpine.XText("'( ' + latestRelease + ' )'"),
		),
	)
}
