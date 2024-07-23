package layout

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func dashboardAside() gomponents.Node {
	return html.Aside(
		components.Classes{
			"w-[80px] h-[100dvh] bg-base-300 shadow-sm p-4": true,
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
				html.Class("text-xs text-center font-bold mt-1"),
				gomponents.Text("PG Back Web"),
			),
		),

		html.Div(
			html.Class("mt-6 space-y-4"),
			dashboardAsideItem(
				lucide.Database,
				"Databases",
				"/dashboard/databases",
			),

			dashboardAsideItem(
				lucide.HardDrive,
				"Destinations",
				"/dashboard/destinations",
			),

			dashboardAsideItem(
				lucide.DatabaseBackup,
				"Backups",
				"/dashboard/backups",
			),

			dashboardAsideItem(
				lucide.List,
				"Executions",
				"/dashboard/executions",
			),

			dashboardAsideItem(
				lucide.User,
				"Profile",
				"/dashboard/profile",
			),

			dashboardAsideItem(
				lucide.Info,
				"About",
				"/dashboard/about",
			),
		),
	)
}

func dashboardAsideItem(
	icon func(children ...gomponents.Node) gomponents.Node,
	text, link string,
) gomponents.Node {
	return html.A(
		html.Class("block flex flex-col items-center justify-center"),
		html.Href(link),
		html.Button(
			html.Class("btn btn-ghost btn-neutral btn-square"),
			icon(html.Class("size-6")),
		),
		html.Span(
			html.Class("text-xs"),
			gomponents.Text(text),
		),
	)
}
