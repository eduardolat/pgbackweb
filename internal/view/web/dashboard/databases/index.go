package databases

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	return echoutil.RenderGomponent(c, http.StatusOK, indexPage())
}

func indexPage() gomponents.Node {
	content := []gomponents.Node{
		html.Div(
			html.Class("flex justify-between items-start"),
			component.H1Text("Databases"),
			createDatabaseButton(),
		),
		html.Div(
			html.Div(
				html.Class("mt-4 overflow-x-auto"),
				html.Table(
					html.Class("table text-nowrap"),
					html.THead(
						html.Tr(
							html.Th(),
							html.Th(component.SpanText("Name")),
							html.Th(component.SpanText("Version")),
							html.Th(component.SpanText("Connection string")),
							html.Th(component.SpanText("Created at")),
						),
					),
					html.TBody(
						htmx.HxGet("/dashboard/databases/list?page=1"),
						htmx.HxTrigger("load"),
						htmx.HxIndicator("#list-databases-loading"),
					),
				),
			),
			html.Div(
				html.Class("flex justify-center mt-4"),
				component.HxLoadingLg("list-databases-loading"),
			),
		),
	}

	return layout.Dashboard(layout.DashboardParams{
		Title: "Databases",
		Body:  content,
	})
}
