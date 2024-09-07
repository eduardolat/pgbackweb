package layout

import (
	"fmt"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
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
			"sticky top-0 z-50":                  true,
		},
		html.Div(
			html.Class("flex justify-start items-center space-x-2"),
			component.ChangeThemeButton(component.ChangeThemeButtonParams{
				Position:    component.DropdownPositionBottom,
				AlignsToEnd: true,
				Size:        component.SizeMd,
			}),
			component.StarOnGithub(component.SizeMd),
			dashboardHeaderCheckForUpdates(),
			component.HxLoadingMd("header-indicator"),
		),
		html.Div(
			html.Class("flex justify-end items-center space-x-2"),
			html.A(
				html.Href("https://discord.gg/BmAwq29UZ8"),
				html.Target("_blank"),
				html.Class("btn btn-ghost btn-neutral"),
				component.SpanText("Chat on Discord"),
				lucide.ExternalLink(),
			),
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

func dashboardHeaderCheckForUpdates() gomponents.Node {
	return html.A(
		alpine.XData("githubRepoInfo"),
		alpine.XCloak(),
		alpine.XShow(fmt.Sprintf(
			"latestRelease !== '' && latestRelease !== '%s'",
			config.Version,
		)),

		components.Classes{
			"btn btn-warning": true,
		},
		html.Href("https://github.com/eduardolat/pgbackweb/releases"),
		html.Target("_blank"),
		lucide.ExternalLink(),
		component.SpanText("Update available"),
		html.Span(
			alpine.XShow("stars"),
			alpine.XText("'( ' + latestRelease + ' )'"),
		),
	)
}
