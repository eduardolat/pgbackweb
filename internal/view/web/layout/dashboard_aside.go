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
			html.Class("block w-full flex justify-center"),
			html.Href("https://github.com/eduardolat/pgbackweb"),
			html.Target("_blank"),
			html.Img(
				html.Src("/images/logo.png"),
				html.Alt("PG Back Web"),
				html.Class("w-[50px] h-auto"),
			),
		),

		html.Div(
			html.Class("mt-8 space-y-4"),
			dashboardAsideItem(
				lucide.Database,
				"Databases",
			),

			dashboardAsideItem(
				lucide.HardDrive,
				"Destinations",
			),

			dashboardAsideItem(
				lucide.DatabaseBackup,
				"Backups",
			),

			dashboardAsideItem(
				lucide.List,
				"Executions",
			),

			dashboardAsideItem(
				lucide.User,
				"Profile",
			),

			dashboardAsideItem(
				lucide.Info,
				"About",
			),
		),
	)
}

func dashboardAsideItem(
	icon func(children ...gomponents.Node) gomponents.Node,
	text string,
) gomponents.Node {
	return html.Div(
		html.Class("flex flex-col items-center justify-center"),
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
