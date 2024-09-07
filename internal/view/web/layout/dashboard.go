package layout

import (
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type DashboardParams struct {
	Title string
	Body  []gomponents.Node
}

func Dashboard(reqCtx reqctx.Ctx, params DashboardParams) gomponents.Node {
	if reqCtx.IsHTMXBoosted {
		return component.RenderableGroup(params.Body)
	}

	title := "PG Back Web"
	if params.Title != "" {
		title = params.Title + " - " + title
	}

	return components.HTML5(components.HTML5Props{
		Language: "en",
		Title:    title,
		Head: []gomponents.Node{
			head(),
		},
		Body: []gomponents.Node{
			htmx.HxIndicator("#header-indicator"),
			components.Classes{
				"w-screen h-screen bg-base-200":      true,
				"flex justify-start overflow-hidden": true,
			},
			dashboardAside(),
			html.Div(
				html.Class("flex-grow overflow-y-auto"),
				dashboardHeader(),
				html.Main(
					html.ID("dashboard-main"),
					html.Class("p-4"),
					gomponents.Group(params.Body),
				),
			),
		},
	})
}
