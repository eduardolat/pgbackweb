package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) manualRunHandler(c echo.Context) error {
	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	go func() {
		_ = h.servs.ExecutionsService.RunExecution(context.Background(), backupID)
	}()

	return respondhtmx.ToastSuccess(c, "Backup started, check the backup executions for more details")
}

func manualRunbutton(backupID uuid.UUID) nodx.Node {
	return component.OptionsDropdownButton(
		htmx.HxPost("/dashboard/backups/"+backupID.String()+"/run"),
		htmx.HxDisabledELT("this"),
		lucide.Zap(),
		component.SpanText("Run backup now"),
	)
}
