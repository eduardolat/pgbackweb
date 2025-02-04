package backups

import (
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmxserver"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) deleteBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return htmxserver.RespondToastError(c, err.Error())
	}

	if err = h.servs.BackupsService.DeleteBackup(ctx, backupID); err != nil {
		return htmxserver.RespondToastError(c, err.Error())
	}

	return htmxserver.RespondRefresh(c)
}

func deleteBackupButton(backupID uuid.UUID) nodx.Node {
	return component.OptionsDropdownButton(
		htmx.HxDelete("/dashboard/backups/"+backupID.String()),
		htmx.HxConfirm("Are you sure you want to delete this backup?"),
		lucide.Trash(),
		component.SpanText("Delete backup"),
	)
}
