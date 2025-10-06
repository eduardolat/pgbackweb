package layout

import (
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
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
				nodx.Src(pathutil.BuildPath("/images/logo.png")),
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
				pathutil.BuildPath("/dashboard"),
				true,
			),

			dashboardAsideItem(
				lucide.Database,
				"Databases",
				pathutil.BuildPath("/dashboard/databases"),
				false,
			),

			dashboardAsideItem(
				lucide.HardDrive,
				"Destinations",
				pathutil.BuildPath("/dashboard/destinations"),
				false,
			),

			dashboardAsideItem(
				lucide.DatabaseBackup,
				"Backup tasks",
				pathutil.BuildPath("/dashboard/backups"),
				false,
			),

			dashboardAsideItem(
				lucide.List,
				"Executions",
				pathutil.BuildPath("/dashboard/executions"),
				false,
			),

			dashboardAsideItem(
				lucide.ArchiveRestore,
				"Restorations",
				pathutil.BuildPath("/dashboard/restorations"),
				false,
			),

			dashboardAsideItem(
				lucide.Webhook,
				"Webhooks",
				pathutil.BuildPath("/dashboard/webhooks"),
				false,
			),

			dashboardAsideItem(
				lucide.User,
				"Profile",
				pathutil.BuildPath("/dashboard/profile"),
				false,
			),

			dashboardAsideItem(
				lucide.Info,
				"About",
				pathutil.BuildPath("/dashboard/about"),
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
