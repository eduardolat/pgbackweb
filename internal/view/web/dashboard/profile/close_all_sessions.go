package profile

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func closeAllSessionsForm(sessions []dbgen.Session) gomponents.Node {
	return component.CardBox(component.CardBoxParams{
		Children: []gomponents.Node{
			component.H2Text("Close all sessions"),
			component.PText("This will log you out from all devices including this one."),
			html.Button(
				htmx.HxPost("/auth/logout-all"),
				htmx.HxDisabledELT("this"),
				htmx.HxConfirm("Are you sure you want to close all your sessions?"),
				html.Class("mt-2 btn btn-error"),
				component.SpanText("Close all sessions"),
				lucide.LogOut(),
			),

			html.Div(html.Class("divider")),

			component.H2Text("Active sessions"),
			component.PText("All sessions are open for a maximum of 12 hours."),
			html.Div(
				html.Class("overflow-x-auto"),
				html.Table(
					html.Class("table"),
					html.THead(
						html.Tr(
							html.Th(component.SpanText("Login time")),
							html.Th(component.SpanText("IP address")),
							html.Th(component.SpanText("User agent")),
						),
					),
					html.TBody(
						component.GMap(sessions, func(session dbgen.Session) gomponents.Node {
							return html.Tr(
								html.Td(component.SpanText(
									session.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
								)),
								html.Td(component.SpanText(session.Ip)),
								html.Td(component.SpanText(session.UserAgent)),
							)
						}),
					),
				),
			),
		},
	})
}
