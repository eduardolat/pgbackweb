package webhooks

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
)

func (h *handlers) duplicateWebhookHandler(c echo.Context) error {
	ctx := c.Request().Context()

	webhookID, err := uuid.Parse(c.Param("webhookID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	if _, err = h.servs.WebhooksService.DuplicateWebhook(ctx, webhookID); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return htmx.RespondRefresh(c)
}

func duplicateWebhookButton(webhookID uuid.UUID) gomponents.Node {
	return component.OptionsDropdownButton(
		htmx.HxPost("/dashboard/webhooks/"+webhookID.String()+"/duplicate"),
		htmx.HxConfirm("Are you sure you want to duplicate this webhook?"),
		lucide.CopyPlus(),
		component.SpanText("Duplicate webhook"),
	)
}
