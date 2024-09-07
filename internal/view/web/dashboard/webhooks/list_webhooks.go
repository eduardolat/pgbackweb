package webhooks

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) listWebhooksHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData struct {
		Page int `query:"page" validate:"required,min=1"`
	}
	if err := c.Bind(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	pagination, whooks, err := h.servs.WebhooksService.PaginateWebhooks(
		ctx, webhooks.PaginateWebhooksParams{
			Page:  formData.Page,
			Limit: 20,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, listWebhooks(pagination, whooks),
	)
}

func listWebhooks(
	pagination paginateutil.PaginateResponse,
	whooks []dbgen.Webhook,
) gomponents.Node {
	if len(whooks) < 1 {
		return component.EmptyResultsTr(component.EmptyResultsParams{
			IsLarge:  true,
			Title:    "No webhooks found",
			Subtitle: "Wait for the first webhook to appear here",
		})
	}

	trs := []gomponents.Node{}
	for _, whook := range whooks {
		trs = append(trs, html.Tr(
			html.Td(
				html.Class("w-[40px]"),
				html.Div(
					html.Class("flex justify-start space-x-1"),
					webhookExecutionsButton(whook.ID),
					runWebhookButton(whook.ID),
					editWebhookButton(whook.ID),
					deleteWebhookButton(whook.ID),
				),
			),
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-2"),
					component.IsActivePing(whook.IsActive),
					component.SpanText(whook.Name),
				),
			),
			html.Td(component.SpanText(
				func() string {
					if name, ok := webhooks.FullEventTypes[whook.EventType]; ok {
						return name
					}
					return whook.EventType
				}(),
			)),
			html.Td(component.SpanText(fmt.Sprintf("%d", len(whook.TargetIds)))),
			html.Td(component.SpanText(
				whook.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, html.Tr(
			htmx.HxGet(func() string {
				url := "/dashboard/webhooks/list"
				url = strutil.AddQueryParamToUrl(url, "page", fmt.Sprintf("%d", pagination.NextPage))
				return url
			}()),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
		))
	}

	return component.RenderableGroup(trs)
}
