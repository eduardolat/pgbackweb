package layout

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type DashboardParams struct {
	Title       string
	Body        []gomponents.Node
	LoadChartJS bool
}

func Dashboard(params DashboardParams) gomponents.Node {
	title := "PG Back Web"
	if params.Title != "" {
		title = params.Title + " - " + title
	}

	return components.HTML5(components.HTML5Props{
		Language: "en",
		Title:    title,
		Head: []gomponents.Node{
			head(headParams{
				LoadChartJS: params.LoadChartJS,
			}),
		},
		Body: []gomponents.Node{
			components.Classes{
				"w-screen h-screen bg-base-200":      true,
				"flex justify-start overflow-hidden": true,
			},
			dashboardAside(),
			html.Div(
				html.Class("flex-grow overflow-y-auto"),
				dashboardHeader(),
				html.Main(
					html.Class("p-4"),
					gomponents.Group(params.Body),
				),
			),
		},
	})
}
