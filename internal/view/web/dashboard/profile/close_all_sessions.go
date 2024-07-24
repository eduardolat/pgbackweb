package profile

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func closeAllSessionsForm() gomponents.Node {
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
		},
	})
}
