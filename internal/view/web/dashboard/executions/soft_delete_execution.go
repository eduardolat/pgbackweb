package executions

import (
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) deleteExecutionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	executionID, err := uuid.Parse(c.Param("executionID"))
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	err = h.servs.ExecutionsService.SoftDeleteExecution(ctx, executionID)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.Refresh(c)
}

func deleteExecutionButton(executionID uuid.UUID) nodx.Node {
	return nodx.Button(
		htmx.HxDelete("/dashboard/executions/"+executionID.String()),
		htmx.HxDisabledELT("this"),
		htmx.HxConfirm("Are you sure you want to delete this execution? It will delete the backup file from the destination and it can't be recovered."),
		nodx.Class("btn btn-error btn-outline"),
		component.SpanText("Delete"),
		lucide.Trash(),
	)
}
