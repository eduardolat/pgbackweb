package webhooks

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) deleteWebhookHandler(c echo.Context) error {
	ctx := c.Request().Context()

	webhookID, err := uuid.Parse(c.Param("webhookID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	if err = h.servs.WebhooksService.DeleteWebhook(ctx, webhookID); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRefresh(c)
}

func deleteWebhookButton(webhookID uuid.UUID) gomponents.Node {
	return html.Div(
		html.Class("inline-block tooltip tooltip-right"),
		html.Data("tip", "Delete webhook"),
		html.Button(
			htmx.HxDelete("/dashboard/webhooks/"+webhookID.String()),
			htmx.HxConfirm("Are you sure you want to delete this webhook?"),
			html.Class("btn btn-error btn-square btn-sm btn-ghost"),
			lucide.Trash(),
		),
	)
}
