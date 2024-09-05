package webhooks

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
	return echoutil.RenderGomponent(
		c, http.StatusOK, indexPage(reqCtx),
	)
}

func indexPage(reqCtx reqctx.Ctx) gomponents.Node {
	content := []gomponents.Node{
		html.Div(
			html.Class("flex justify-between items-start"),
			component.H1Text("Webhooks"),
			createWebhookButton(),
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
								html.Th(component.SpanText("Event type")),
								html.Th(component.SpanText("Targets")),
								html.Th(component.SpanText("Created at")),
							),
						),
						html.TBody(
							htmx.HxGet("/dashboard/webhooks/list?page=1"),
							htmx.HxTrigger("load"),
							htmx.HxIndicator("#list-webhooks-loading"),
						),
					),

					html.Div(
						html.Class("flex justify-center mt-4"),
						component.HxLoadingLg("list-webhooks-loading"),
					),
				),
			},
		}),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: "Webhooks",
		Body:  content,
	})
}
