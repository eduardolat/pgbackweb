package webhooks

import (
	"context"
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) runWebhookHandler(c echo.Context) error {
	webhookID, err := uuid.Parse(c.Param("webhookID"))
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
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
		err = h.servs.WebhooksService.SendWebhookRequest(ctx, webhook, webhooks.WebhookPayload{
			EventType: "test_webhook",
			Msg:       "Webhook test event"})
		if err != nil {
			logger.Error("error sending webhook request", logger.KV{
				"webhook_id": webhook.ID,
				"error":      err.Error(),
			})
		}
	}()

	return respondhtmx.ToastSuccess(c, "Running webhook, check the webhook executions for more details")
}

func runWebhookButton(webhookID uuid.UUID) nodx.Node {
	return component.OptionsDropdownButton(
		htmx.HxPost(pathutil.BuildPath(fmt.Sprintf("/dashboard/webhooks/%s/run", webhookID))),
		htmx.HxDisabledELT("this"),
		lucide.Zap(),
		component.SpanText("Run webhook now"),
	)
}
