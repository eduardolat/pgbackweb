package databases

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	reqCtx := reqctx.GetCtx(c)
	return echoutil.RenderGomponent(c, http.StatusOK, indexPage(reqCtx))
}

func indexPage(reqCtx reqctx.Ctx) gomponents.Node {
	content := []gomponents.Node{
		html.Div(
			html.Class("flex justify-between items-start"),
			component.H1Text("Databases"),
			createDatabaseButton(),
		),
		component.CardBox(component.CardBoxParams{
			Class: "mt-4",
			Children: []gomponents.Node{
				html.Div(
					html.Class("overflow-x-auto"),
					html.Table(
						html.Class("table text-nowrap"),
						html.THead(
							html.Tr(
								html.Th(component.SpanText("Actions")),
								html.Th(component.SpanText("Name")),
								html.Th(component.SpanText("Version")),
								html.Th(component.SpanText("Connection string")),
								html.Th(component.SpanText("Created at")),
							),
						),
						html.TBody(
							component.SkeletonTr(8),
							htmx.HxGet("/dashboard/databases/list?page=1"),
							htmx.HxTrigger("load"),
						),
					),
				),
			},
		}),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: "Databases",
		Body:  content,
	})
}
