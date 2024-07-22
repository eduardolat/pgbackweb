package layout

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type AuthParams struct {
	Title string
	Body  []gomponents.Node
}

func Auth(params AuthParams) gomponents.Node {
	title := "PG Back Web"
	if params.Title != "" {
		title = params.Title + " - " + title
	}

	return components.HTML5(components.HTML5Props{
		Language: "en",
		Title:    title,
		Head: []gomponents.Node{
			html.Link(
				html.Rel("shortcut icon"),
				html.Href("/favicon.ico"),
			),

			html.Link(
				html.Rel("stylesheet"),
				html.Href("/css/style.css"),
			),

			html.Script(
				html.Src("/js/alpinejs-3.14.1.min.js"),
				html.Defer(),
			),
			html.Script(
				html.Src("/js/htmx-2.0.1.min.js"),
				html.Defer(),
			),
			html.Script(
				html.Src("/js/theme-change-2.0.2.min.js"),
			),
		},
		Body: []gomponents.Node{
			components.Classes{
				"w-screen h-screen p-4":               true,
				"grid grid-cols-1 place-items-center": true,
				"bg-base-300 overflow-y-auto":         true,
			},
			html.Main(
				components.Classes{
					"w-full max-w-[600px]":  true,
					"rounded-box shadow-md": true,
					"bg-base-100 p-4":       true,
				},
				gomponents.Group(params.Body),
			),
		},
	})
}
