package layout

import (
	"fmt"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func dashboardAside() gomponents.Node {
	return html.Aside(
		html.ID("dashboard-aside"),
		components.Classes{
			"flex-none h-[100dvh] bg-base-300 shadow-sm p-4": true,
			"overflow-y-auto overflow-x-hidden":              true,
		},

		html.A(
			html.Class("block flex flex-col justify-center items-center"),
			html.Href("https://github.com/eduardolat/pgbackweb"),
			html.Target("_blank"),
			html.Img(
				html.Src("/images/logo.png"),
				html.Alt("PG Back Web"),
				html.Class("w-[50px] h-auto"),
			),
			html.Span(
				html.Class("text-xs text-nowrap text-center font-bold mt-1"),
				html.Span(
					html.Class("block"),
					gomponents.Text("PG Back"),
				),
				html.Span(
					html.Class("block"),
					gomponents.Text("Web"),
				),
			),
		),

		html.Div(
			html.Class("mt-6 space-y-4"),

			dashboardAsideItem(
				lucide.LayoutDashboard,
				"Summary",
				"/dashboard",
				true,
			),

			dashboardAsideItem(
				lucide.Database,
				"Databases",
				"/dashboard/databases",
				false,
			),

			dashboardAsideItem(
				lucide.HardDrive,
				"Destinations",
				"/dashboard/destinations",
				false,
			),

			dashboardAsideItem(
				lucide.DatabaseBackup,
				"Backups",
				"/dashboard/backups",
				false,
			),

			dashboardAsideItem(
				lucide.List,
				"Executions",
				"/dashboard/executions",
				false,
			),

			dashboardAsideItem(
				lucide.ArchiveRestore,
				"Restorations",
				"/dashboard/restorations",
				false,
			),

			dashboardAsideItem(
				lucide.Webhook,
				"Webhooks",
				"/dashboard/webhooks",
				false,
			),

			dashboardAsideItem(
				lucide.User,
				"Profile",
				"/dashboard/profile",
				false,
			),

			dashboardAsideItem(
				lucide.Info,
				"About",
				"/dashboard/about",
				false,
			),
		),
	)
}

func dashboardAsideItem(
	icon func(children ...gomponents.Node) gomponents.Node,
	text, link string, strict bool,
) gomponents.Node {
	return html.A(
		alpine.XData(fmt.Sprintf("dashboardAsideItem('%s', %t)", link, strict)),
		html.Class("block flex flex-col items-center justify-center group"),
		html.Href(link),
		html.Button(
			alpine.XBind("class", `{'btn-active': is_active}`),
			html.Class("btn btn-ghost btn-neutral btn-square group-hover:btn-active"),
			icon(html.Class("size-6")),
		),
		html.Span(
			html.Class("text-xs"),
			gomponents.Text(text),
		),
	)
}
