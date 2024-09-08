package backups

import (
	"context"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
)

func (h *handlers) manualRunHandler(c echo.Context) error {
	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	go func() {
		_ = h.servs.ExecutionsService.RunExecution(context.Background(), backupID)
	}()

	return htmx.RespondToastSuccess(c, "Backup started, check the backup executions for more details")
}

func manualRunbutton(backupID uuid.UUID) gomponents.Node {
	return component.OptionsDropdownButton(
		htmx.HxPost("/dashboard/backups/"+backupID.String()+"/run"),
		htmx.HxDisabledELT("this"),
		lucide.Zap(),
		component.SpanText("Run backup now"),
	)
}
