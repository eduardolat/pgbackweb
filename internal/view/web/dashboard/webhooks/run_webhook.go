package webhooks

import (
	"context"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
)

func (h *handlers) runWebhookHandler(c echo.Context) error {
	webhookID, err := uuid.Parse(c.Param("webhookID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	go func() {
		ctx := context.Background()
		webhook, err := h.servs.WebhooksService.GetWebhook(ctx, webhookID)
		if err != nil {
			logger.Error("error getting webhook", logger.KV{
				"webhook_id": webhookID.String(),
				"error":      err.Error(),
			})
			return
		}
		err = h.servs.WebhooksService.SendWebhookRequest(ctx, webhook)
		if err != nil {
			logger.Error("error sending webhook request", logger.KV{
				"webhook_id": webhook.ID,
				"error":      err.Error(),
			})
		}
	}()

	return htmx.RespondToastSuccess(c, "Running webhook, check the webhook executions for more details")
}

func runWebhookButton(webhookID uuid.UUID) gomponents.Node {
	return component.OptionsDropdownButton(
		htmx.HxPost("/dashboard/webhooks/"+webhookID.String()+"/run"),
		htmx.HxDisabledELT("this"),
		lucide.Zap(),
		component.SpanText("Run webhook now"),
	)
}
