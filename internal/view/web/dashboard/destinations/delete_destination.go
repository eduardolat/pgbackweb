package destinations

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) deleteDestinationHandler(c echo.Context) error {
	ctx := c.Request().Context()

	destinationID, err := uuid.Parse(c.Param("destinationID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	err = h.servs.DestinationsService.DeleteDestination(ctx, destinationID)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRefresh(c)
}

func deleteDestinationButton(destinationID uuid.UUID) gomponents.Node {
	return html.Div(
		html.Class("inline-block tooltip tooltip-right"),
		html.Data("tip", "Delete destination"),
		html.Button(
			htmx.HxDelete("/dashboard/destinations/"+destinationID.String()),
			htmx.HxConfirm("Are you sure you want to delete this destination?"),
			html.Class("btn btn-error btn-square btn-sm btn-ghost"),
			lucide.Trash(),
		),
	)
}
