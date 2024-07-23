package layout

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type DashboardParams struct {
	Title string
	Body  []gomponents.Node
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
			html.Link(html.Rel("shortcut icon"), html.Href("/favicon.ico")),
			html.Link(html.Rel("stylesheet"), html.Href("/css/style.css")),
			html.Script(gomponents.Attr("type", "module"), html.Src("/js/app.js")),

			html.Script(html.Src("/libs/htmx/htmx-2.0.1.min.js"), html.Defer()),
			html.Script(html.Src("/libs/alpinejs/alpinejs-3.14.1.min.js"), html.Defer()),
			html.Script(html.Src("/libs/theme-change/theme-change-2.0.2.min.js")),

			html.Link(html.Rel("stylesheet"), html.Href("/libs/notyf/notyf-3.10.0.min.css")),
			html.Script(html.Src("/libs/notyf/notyf-3.10.0.min.js")),
		},
		Body: []gomponents.Node{
			components.Classes{
				"w-screen h-screen bg-base-200": true,
				"flex justify-start":            true,
			},
			dashboardAside(),
			html.Div(
				html.Class("flex-grow"),
				dashboardHeader(),
				html.Main(
					html.Class("p-4 overflow-auto"),
					gomponents.Group(params.Body),
				),
			),
		},
	})
}
