package layout

import (
	"fmt"

	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func dashboardAside() nodx.Node {
	return nodx.Aside(
		nodx.Id("dashboard-aside"),
		nodx.ClassMap{
			"flex-none h-[100dvh] bg-base-300 shadow-sm p-4": true,
			"overflow-y-auto overflow-x-hidden":              true,
		},

		nodx.A(
			nodx.Class("block flex flex-col justify-center items-center"),
			nodx.Href("https://github.com/eduardolat/pgbackweb"),
			nodx.Target("_blank"),
			nodx.Img(
				nodx.Src("/images/logo.png"),
				nodx.Alt("PG Back Web"),
				nodx.Class("w-[50px] h-auto"),
			),
			nodx.SpanEl(
				nodx.Class("text-xs text-nowrap text-center font-bold mt-1"),
				nodx.SpanEl(
					nodx.Class("block"),
					nodx.Text("PG Back"),
				),
				nodx.SpanEl(
					nodx.Class("block"),
					nodx.Text("Web"),
				),
			),
		),

		nodx.Div(
			nodx.Class("mt-6 space-y-4"),

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
	icon func(children ...nodx.Node) nodx.Node,
	text, link string, strict bool,
) nodx.Node {
	return nodx.A(
		alpine.XData(fmt.Sprintf("alpineDashboardAsideItem('%s', %t)", link, strict)),
		nodx.Class("block flex flex-col items-center justify-center group"),

		nodx.Href(link),
		htmx.HxBoost("true"),
		htmx.HxTarget("#dashboard-main"),
		htmx.HxSwap("transition:true show:unset"),

		nodx.Button(
			alpine.XBind("class", `{'btn-active': is_active}`),
			nodx.Class("btn btn-ghost btn-neutral btn-square group-hover:btn-active"),
			icon(nodx.Class("size-6")),
		),
		nodx.SpanEl(
			nodx.Class("text-xs"),
			nodx.Text(text),
		),
	)
}
