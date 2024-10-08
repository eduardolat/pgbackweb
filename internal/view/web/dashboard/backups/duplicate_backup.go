package backups

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
)

func (h *handlers) duplicateBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	if _, err = h.servs.BackupsService.DuplicateBackup(ctx, backupID); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRefresh(c)
}

func duplicateBackupButton(backupID uuid.UUID) gomponents.Node {
	return component.OptionsDropdownButton(
		htmx.HxPost("/dashboard/backups/"+backupID.String()+"/duplicate"),
		htmx.HxConfirm("Are you sure you want to duplicate this backup?"),
		lucide.CopyPlus(),
		component.SpanText("Duplicate backup"),
	)
}
