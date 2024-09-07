package dashboard

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func healthButtonHandler(servs *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		databasesQty, err := servs.DatabasesService.GetDatabasesQty(ctx)
		if err != nil {
			return htmx.RespondToastError(c, err.Error())
		}
		destinationsQty, err := servs.DestinationsService.GetDestinationsQty(ctx)
		if err != nil {
			return htmx.RespondToastError(c, err.Error())
		}

		return echoutil.RenderGomponent(c, http.StatusOK, healthButton(
			databasesQty, destinationsQty,
		))
	}
}

func healthButton(
	databasesQty dbgen.DatabasesServiceGetDatabasesQtyRow,
	destinationsQty dbgen.DestinationsServiceGetDestinationsQtyRow,
) gomponents.Node {
	isHealthy := true

	if databasesQty.Unhealthy > 0 {
		isHealthy = false
	}
	if destinationsQty.Unhealthy > 0 {
		isHealthy = false
	}

	pingColor := component.ColorSuccess
	if !isHealthy {
		pingColor = component.ColorError
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Health status",
		Content: []gomponents.Node{
			component.PText(`
				The health check for both databases and destinations runs automatically
				every 10 minutes, when PG Back Web starts, and when you click the
				"Test connection" button on each resource. You can see additional
				information and error messages by clicking the health check button
				for each resource.
			`),
			html.Table(
				html.Class("table mt-2"),
				html.THead(
					html.Tr(
						html.Th(component.SpanText("Resource")),
						html.Th(component.SpanText("Total")),
						html.Th(component.SpanText("Healthy")),
						html.Th(component.SpanText("Unhealthy")),
					),
				),
				html.TBody(
					html.Tr(
						html.Td(component.SpanText("Databases")),
						html.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.All))),
						html.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.Healthy))),
						html.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.Unhealthy))),
					),
					html.Tr(
						html.Td(component.SpanText("Destinations")),
						html.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.All))),
						html.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.Healthy))),
						html.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.Unhealthy))),
					),
				),
			),
		},
	})

	return html.Div(
		html.Class("inline-block"),
		mo.HTML,
		html.Button(
			mo.OpenerAttr,
			html.Class("btn btn-ghost btn-neutral"),
			component.SpanText("Health status"),
			component.Ping(pingColor),
		),
	)
}
