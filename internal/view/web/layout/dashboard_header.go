package layout

import (
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func dashboardHeader() nodx.Node {
	return nodx.Header(
		nodx.ClassMap{
			"sticky top-0 z-50":                 true,
			"space-x-4 p-4 min-w-max":           true,
			"w-[full] bg-base-200 shadow-sm":    true,
			"flex items-center justify-between": true,
		},
		nodx.Div(
			nodx.Class("flex justify-start items-center space-x-2"),
			component.SupportProjectButton(component.SizeSm),
			component.ChangeThemeButton(component.ChangeThemeButtonParams{
				Position: component.DropdownPositionBottom,
				Size:     component.SizeSm,
			}),
			component.StarOnGithub(component.SizeSm),
			dashboardHeaderUpdates(),
		),
		nodx.Div(
			nodx.Class("flex justify-end items-center space-x-2"),
			nodx.Div(
				htmx.HxGet(pathutil.BuildPath("/dashboard/health-button")),
				htmx.HxSwap("outerHTML"),
				htmx.HxTrigger("load once"),
			),
			nodx.A(
				nodx.Href("https://ufobackup.uforg.dev/r/community"),
				nodx.Target("_blank"),
				nodx.Class("btn btn-ghost btn-neutral"),
				component.SpanText("Join the community"),
				lucide.ExternalLink(),
			),
			nodx.Button(
				htmx.HxPost(pathutil.BuildPath("/auth/logout")),
				htmx.HxDisabledELT("this"),
				nodx.Class("btn btn-ghost btn-neutral"),
				component.SpanText("Log out"),
				lucide.LogOut(),
			),
		),
	)
}
