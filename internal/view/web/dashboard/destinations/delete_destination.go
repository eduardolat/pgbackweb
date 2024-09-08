package destinations

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
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
	return component.OptionsDropdownButton(
		htmx.HxDelete("/dashboard/destinations/"+destinationID.String()),
		htmx.HxConfirm("Are you sure you want to delete this destination?"),
		lucide.Trash(),
		component.SpanText("Delete destination"),
	)
}
