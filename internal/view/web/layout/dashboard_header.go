package layout

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func dashboardHeader() gomponents.Node {
	return html.Header(
		components.Classes{
			"w-[full] bg-base-200 p-4 shadow-sm": true,
			"flex items-center justify-between":  true,
			"sticky top-0":                       true,
		},
		html.Div(
			html.Class("flex justify-start items-center space-x-2"),
			component.ChangeThemeButton(component.ChangeThemeButtonParams{
				Position:    component.DropdownPositionBottom,
				AlignsToEnd: true,
				Size:        component.SizeMd,
			}),
			component.StarOnGithub(component.SizeMd),
		),
		html.Div(
			html.Button(
				htmx.HxPost("/auth/logout"),
				htmx.HxDisabledELT("this"),
				html.Class("btn btn-ghost btn-neutral"),
				component.SpanText("Log out"),
				lucide.LogOut(),
			),
		),
	)
}
