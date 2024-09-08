package backups

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
)

func (h *handlers) deleteBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	if err = h.servs.BackupsService.DeleteBackup(ctx, backupID); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRefresh(c)
}

func deleteBackupButton(backupID uuid.UUID) gomponents.Node {
	return component.OptionsDropdownButton(
		htmx.HxDelete("/dashboard/backups/"+backupID.String()),
		htmx.HxConfirm("Are you sure you want to delete this backup?"),
		lucide.Trash(),
		component.SpanText("Delete backup"),
	)
}
