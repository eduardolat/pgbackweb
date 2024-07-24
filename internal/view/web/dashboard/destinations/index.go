package destinations

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
			component.H1Text("Destinations"),
			createDestinationButton(),
		),
		html.Div(
			html.Div(
				html.Class("mt-4 overflow-x-auto"),
				html.Table(
					html.Class("table"),
					html.THead(
						html.Tr(
							html.Th(),
							html.Th(component.SpanText("Name")),
							html.Th(component.SpanText("Bucket name")),
							html.Th(component.SpanText("Endpoint")),
							html.Th(component.SpanText("Region")),
							html.Th(component.SpanText("Access key")),
							html.Th(component.SpanText("Secret key")),
							html.Th(component.SpanText("Created at")),
						),
					),
					html.TBody(
						htmx.HxGet("/dashboard/destinations/list?page=1"),
						htmx.HxTrigger("load"),
						htmx.HxIndicator("#list-destinations-loading"),
					),
				),
			),
			html.Div(
				html.Class("flex justify-center mt-4"),
				component.HxLoadingLg("list-destinations-loading"),
			),
		),
	}

	return layout.Dashboard(layout.DashboardParams{
		Title: "Destinations",
		Body:  content,
	})
}
