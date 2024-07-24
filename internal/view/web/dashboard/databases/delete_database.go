package databases

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
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
	return html.Div(
		html.Class("inline-block tooltip tooltip-right"),
		html.Data("tip", "Delete database"),
		html.Button(
			htmx.HxDelete("/dashboard/databases/"+databaseID.String()),
			htmx.HxConfirm("Are you sure you want to delete this database?"),
			html.Class("btn btn-error btn-square btn-sm btn-ghost"),
			lucide.Trash(),
		),
	)
}
