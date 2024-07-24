package executions

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) deleteExecutionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	executionID, err := uuid.Parse(c.Param("executionID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	err = h.servs.ExecutionsService.SoftDeleteExecution(ctx, executionID)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRefresh(c)
}

func deleteExecutionButton(executionID uuid.UUID) gomponents.Node {
	return html.Button(
		htmx.HxDelete("/dashboard/executions/"+executionID.String()),
		htmx.HxDisabledELT("this"),
		htmx.HxConfirm("Are you sure you want to delete this execution? It will delete the backup file from the destination and it can't be recovered."),
		html.Class("btn btn-error btn-outline"),
		component.SpanText("Delete"),
		lucide.Trash(),
	)
}
