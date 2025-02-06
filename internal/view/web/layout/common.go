package layout

import (
	"github.com/eduardolat/pgbackweb/internal/view/static"
	nodx "github.com/nodxdev/nodxgo"
)

func commonHtmlDoc(title string, bodyContentGroup nodx.Node) nodx.Node {
	return nodx.Group(
		nodx.DocType(),
		nodx.Html(
			nodx.Lang("en"),
			nodx.Head(
				nodx.TitleEl(nodx.Text(title)),
				commonHead(),
			),
			nodx.Body(bodyContentGroup),
		),
	)
}

func commonHead() nodx.Node {
	href := func(path string) nodx.Node {
		return nodx.Href(static.GetVersionedFilePath(path))
	}

	src := func(path string) nodx.Node {
		return nodx.Src(static.GetVersionedFilePath(path))
	}

	return nodx.Group(
		nodx.Meta(nodx.Charset("utf-8")),
		nodx.Meta(nodx.Name("viewport"), nodx.Content("width=device-width, initial-scale=1")),

		// https://htmx.org/quirks/
		nodx.Meta(nodx.Name("htmx-config"), nodx.Content(`{"disableInheritance":true, "responseHandling": [{"code":"...", "swap": true}]}`)),

		nodx.Link(nodx.Rel("shortcut icon"), href("/favicon.ico")),
		nodx.Link(nodx.Rel("stylesheet"), href("/build/style.min.css")),
		nodx.Script(src("/build/app.min.js")),

		nodx.Script(src("/libs/htmx/htmx-2.0.1.min.js"), nodx.Defer("")),
		nodx.Script(src("/libs/alpinejs/alpinejs-3.14.1.min.js"), nodx.Defer("")),
		nodx.Script(src("/libs/sweetalert2/sweetalert2-11.13.1.min.js")),
		nodx.Script(src("/libs/chartjs/chartjs-4.4.3.umd.min.js")),

		nodx.Link(nodx.Rel("stylesheet"), href("/libs/notyf/notyf-3.10.0.min.css")),
		nodx.Script(src("/libs/notyf/notyf-3.10.0.min.js")),

		nodx.Link(nodx.Rel("stylesheet"), href("/libs/slim-select/slimselect-2.8.2.css")),
		nodx.Script(src("/libs/slim-select/slimselect-2.8.2.min.js")),
	)
}
