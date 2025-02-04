package backups

import (
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmxs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) duplicateBackupHandler(c echo.Context) error {
	ctx := c.Request().Context()

	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	if _, err = h.servs.BackupsService.DuplicateBackup(ctx, backupID); err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	return htmxs.RespondRefresh(c)
}

func duplicateBackupButton(backupID uuid.UUID) nodx.Node {
	return component.OptionsDropdownButton(
		htmx.HxPost("/dashboard/backups/"+backupID.String()+"/duplicate"),
		htmx.HxConfirm("Are you sure you want to duplicate this backup?"),
		lucide.CopyPlus(),
		component.SpanText("Duplicate backup"),
	)
}
