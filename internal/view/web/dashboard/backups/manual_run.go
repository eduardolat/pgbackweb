package backups

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) manualRunHandler(c echo.Context) error {
	ctx := c.Request().Context()

	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	if err := h.servs.ExecutionsService.RunExecution(ctx, backupID); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondToastSuccess(c, "Backup run successfully, check the backup executions for more details")
}

func manualRunbutton(backupID uuid.UUID) gomponents.Node {
	return html.Div(
		html.Class("inline-block tooltip tooltip-right"),
		html.Data("tip", "Run backup now"),
		html.Button(
			htmx.HxPost("/dashboard/backups/"+backupID.String()+"/run"),
			htmx.HxDisabledELT("this"),
			html.Class("btn btn-sm btn-ghost btn-square"),
			lucide.Zap(),
		),
	)
}
