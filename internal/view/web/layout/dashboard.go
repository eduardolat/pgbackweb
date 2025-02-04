package layout

import (
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
)

type DashboardParams struct {
	Title string
	Body  []nodx.Node
}

func Dashboard(reqCtx reqctx.Ctx, params DashboardParams) nodx.Node {
	title := "PG Back Web"
	if params.Title != "" {
		title = params.Title + " - " + title
	}

	if reqCtx.IsHTMXBoosted {
		body := append(params.Body, nodx.TitleEl(nodx.Text(title)))
		return component.RenderableGroup(body)
	}

	body := nodx.Group(
		nodx.ClassMap{
			"w-screen h-screen bg-base-200":      true,
			"flex justify-start overflow-hidden": true,
		},
		dashboardAside(),
		nodx.Div(
			nodx.Class("flex-grow overflow-y-auto"),
			dashboardHeader(),
			nodx.Main(
				nodx.Id("dashboard-main"),
				nodx.Class("p-4"),
				nodx.Group(params.Body...),
			),
		),
	)

	return commonHtmlDoc(title, body)
}
