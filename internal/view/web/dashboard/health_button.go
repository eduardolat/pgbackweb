package dashboard

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
)

func healthButtonHandler(servs *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		databasesQty, err := servs.DatabasesService.GetDatabasesQty(ctx)
		if err != nil {
			return respondhtmx.ToastError(c, err.Error())
		}
		destinationsQty, err := servs.DestinationsService.GetDestinationsQty(ctx)
		if err != nil {
			return respondhtmx.ToastError(c, err.Error())
		}

		return echoutil.RenderNodx(c, http.StatusOK, healthButton(
			databasesQty, destinationsQty,
		))
	}
}

func healthButton(
	databasesQty dbgen.DatabasesServiceGetDatabasesQtyRow,
	destinationsQty dbgen.DestinationsServiceGetDestinationsQtyRow,
) nodx.Node {
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
		Content: []nodx.Node{
			component.PText(`
				The health check for both databases and destinations runs automatically
				every 10 minutes, when PG Back Web starts, and when you click the
				"Test connection" button on each resource. You can see additional
				information and error messages by clicking the health check button
				for each resource.
			`),
			nodx.Table(
				nodx.Class("table mt-2"),
				nodx.Thead(
					nodx.Tr(
						nodx.Th(component.SpanText("Resource")),
						nodx.Th(component.SpanText("Total")),
						nodx.Th(component.SpanText("Healthy")),
						nodx.Th(component.SpanText("Unhealthy")),
					),
				),
				nodx.Tbody(
					nodx.Tr(
						nodx.Td(component.SpanText("Databases")),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.All))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.Healthy))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.Unhealthy))),
					),
					nodx.Tr(
						nodx.Td(component.SpanText("Destinations")),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.All))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.Healthy))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.Unhealthy))),
					),
				),
			),
		},
	})

	return nodx.Div(
		nodx.Class("inline-block"),
		mo.HTML,
		nodx.Button(
			mo.OpenerAttr,
			nodx.Class("btn btn-ghost btn-neutral"),
			component.SpanText("Health status"),
			component.Ping(pingColor),
		),
	)
}
