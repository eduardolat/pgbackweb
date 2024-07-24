package summary

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	ctx := c.Request().Context()

	databasesQty, err := h.servs.DatabasesService.GetDatabasesQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	destinationsQty, err := h.servs.DestinationsService.GetDestinationsQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	backupsQty, err := h.servs.BackupsService.GetBackupsQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	executionsQty, err := h.servs.ExecutionsService.GetExecutionsQty(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK,
		indexPage(databasesQty, destinationsQty, backupsQty, executionsQty),
	)
}

func indexPage(
	databasesQty, destinationsQty, backupsQty, executionsQty int64,
) gomponents.Node {
	countCard := func(title string, count int64) gomponents.Node {
		return component.CardBox(component.CardBoxParams{
			Class: "text-center",
			Children: []gomponents.Node{
				component.H2Text(title),
				html.Span(
					html.Class("text-5xl font-bold"),
					gomponents.Textf("%d", count),
				),
			},
		})
	}

	content := []gomponents.Node{
		component.H1Text("Summary"),
		html.Div(
			html.Class("mt-4 grid grid-cols-4 gap-4"),
			countCard("Databases", databasesQty),
			countCard("Destinations", destinationsQty),
			countCard("Backups", backupsQty),
			countCard("Executions", executionsQty),
		),
	}

	return layout.Dashboard(layout.DashboardParams{
		Title: "Summary",
		Body:  content,
	})
}
