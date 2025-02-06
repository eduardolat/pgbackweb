package profile

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func closeAllSessionsForm(sessions []dbgen.Session) nodx.Node {
	return component.CardBox(component.CardBoxParams{
		Children: []nodx.Node{
			component.H2Text("Close all sessions"),
			component.PText("This will log you out from all devices including this one."),
			nodx.Button(
				htmx.HxPost("/auth/logout-all"),
				htmx.HxDisabledELT("this"),
				htmx.HxConfirm("Are you sure you want to close all your sessions?"),
				nodx.Class("mt-2 btn btn-error"),
				component.SpanText("Close all sessions"),
				lucide.LogOut(),
			),

			nodx.Div(nodx.Class("divider")),

			component.H2Text("Active sessions"),
			component.PText("All sessions are open for a maximum of 12 hours."),
			nodx.Div(
				nodx.Class("overflow-x-auto"),
				nodx.Table(
					nodx.Class("table"),
					nodx.Thead(
						nodx.Tr(
							nodx.Th(component.SpanText("Login time")),
							nodx.Th(component.SpanText("IP address")),
							nodx.Th(component.SpanText("User agent")),
						),
					),
					nodx.Tbody(
						nodx.Map(sessions, func(session dbgen.Session) nodx.Node {
							return nodx.Tr(
								nodx.Td(component.SpanText(
									session.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
								)),
								nodx.Td(component.SpanText(session.Ip)),
								nodx.Td(component.SpanText(session.UserAgent)),
							)
						}),
					),
				),
			),
		},
	})
}
