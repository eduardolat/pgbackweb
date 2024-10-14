package layout

import (
	"github.com/eduardolat/pgbackweb/internal/view/static"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func head() gomponents.Node {
	href := func(path string) gomponents.Node {
		return html.Href(static.GetVersionedFilePath(path))
	}

	src := func(path string) gomponents.Node {
		return html.Src(static.GetVersionedFilePath(path))
	}

	return gomponents.Group([]gomponents.Node{
		html.Link(html.Rel("shortcut icon"), href("/favicon.ico")),
		html.Link(html.Rel("stylesheet"), href("/build/style.min.css")),
		html.Script(src("/build/app.min.js")),

		html.Script(src("/libs/htmx/htmx-2.0.1.min.js"), html.Defer()),
		html.Script(src("/libs/alpinejs/alpinejs-3.14.1.min.js"), html.Defer()),
		html.Script(src("/libs/sweetalert2/sweetalert2-11.13.1.min.js")),
		html.Script(src("/libs/chartjs/chartjs-4.4.3.umd.min.js")),

		html.Link(html.Rel("stylesheet"), href("/libs/notyf/notyf-3.10.0.min.css")),
		html.Script(src("/libs/notyf/notyf-3.10.0.min.js")),

		html.Link(html.Rel("stylesheet"), href("/libs/slim-select/slimselect-2.8.2.css")),
		html.Script(src("/libs/slim-select/slimselect-2.8.2.min.js")),
	})
}
