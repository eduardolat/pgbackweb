package databases

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
)

func (h *handlers) deleteDatabaseHandler(c echo.Context) error {
	ctx := c.Request().Context()

	databaseID, err := uuid.Parse(c.Param("databaseID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	if err = h.servs.DatabasesService.DeleteDatabase(ctx, databaseID); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRefresh(c)
}

func deleteDatabaseButton(databaseID uuid.UUID) gomponents.Node {
	return component.OptionsDropdownButton(
		htmx.HxDelete("/dashboard/databases/"+databaseID.String()),
		htmx.HxConfirm("Are you sure you want to delete this database?"),
		lucide.Trash(),
		component.SpanText("Delete database"),
	)
}
