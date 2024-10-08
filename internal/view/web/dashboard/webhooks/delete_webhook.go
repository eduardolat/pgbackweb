package webhooks

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
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
	return component.OptionsDropdownButton(
		htmx.HxDelete("/dashboard/webhooks/"+webhookID.String()),
		htmx.HxConfirm("Are you sure you want to delete this webhook?"),
		lucide.Trash(),
		component.SpanText("Delete webhook"),
	)
}
