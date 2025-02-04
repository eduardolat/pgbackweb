package webhooks

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmxs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) paginateWebhookExecutionsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	webhookID, err := uuid.Parse(c.Param("webhookID"))
	if err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	var queryData struct {
		Page int `query:"page" validate:"required,min=1"`
	}
	if err := c.Bind(&queryData); err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&queryData); err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	pagination, execs, err := h.servs.WebhooksService.PaginateWebhookExecutions(
		ctx, webhooks.PaginateWebhookExecutionsParams{
			WebhookID: webhookID,
			Page:      queryData.Page,
			Limit:     20,
		},
	)
	if err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	return echoutil.RenderNodx(
		c, http.StatusOK, webhookExecutionsList(webhookID, pagination, execs),
	)
}

func webhookExecutionsList(
	webhookID uuid.UUID,
	pagination paginateutil.PaginateResponse,
	execs []dbgen.WebhookExecution,
) nodx.Node {
	if len(execs) == 0 {
		return component.EmptyResultsTr(component.EmptyResultsParams{
			Title:    "No executions found",
			Subtitle: "Wait for the first execution to appear here",
		})
	}

	trs := []nodx.Node{}
	for _, exec := range execs {
		durationMillis := exec.ResDuration.Int32
		duration := time.Duration(durationMillis) * time.Millisecond

		trs = append(trs, nodx.Tr(
			nodx.Td(
				webhookExecutionDetailsButton(exec, duration),
			),
			nodx.Td(component.SpanText(fmt.Sprintf("%d", exec.ResStatus.Int16))),
			nodx.Td(component.SpanText(exec.ReqMethod.String)),
			nodx.Td(component.SpanText(duration.String())),
			nodx.Td(component.SpanText(
				exec.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, nodx.Tr(
			htmx.HxGet(func() string {
				url := "/dashboard/webhooks/" + webhookID.String() + "/executions"
				url = strutil.AddQueryParamToUrl(url, "page", fmt.Sprintf("%d", pagination.NextPage))
				return url
			}()),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
			htmx.HxIndicator("#webhook-executions-loading"),
		))
	}

	return component.RenderableGroup(trs)
}

func webhookExecutionDetailsButton(
	exec dbgen.WebhookExecution,
	duration time.Duration,
) nodx.Node {
	mo := component.Modal(component.ModalParams{
		Title: "Webhook execution details",
		Content: []nodx.Node{
			nodx.Div(
				nodx.Class("space-y-4"),

				alpine.XData(`{
					processTextareas() {
						const els = [
							$refs.reqHeadersTextarea,
							$refs.reqBodyTextarea,
							$refs.resHeadersTextarea,
							$refs.resBodyTextarea
						]

						for (const el of els) {
							el.value = formatJson(el.value)
						}
					}
				}`),
				alpine.XOn("mouseenter.once", "processTextareas()"),

				nodx.Table(
					nodx.Class("table [&_th]:text-nowrap"),
					nodx.Tr(
						nodx.Td(
							nodx.Colspan("100%"),
							component.H3Text("General"),
						),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("ID")),
						nodx.Td(component.SpanText(exec.ID.String())),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Date")),
						nodx.Td(component.SpanText(
							exec.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
						)),
					),
				),

				nodx.Table(
					nodx.Class("table [&_th]:text-nowrap"),
					nodx.Tr(
						nodx.Td(
							nodx.Colspan("100%"),
							component.H3Text("Request"),
						),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Method")),
						nodx.Td(component.SpanText(exec.ReqMethod.String)),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Headers")),
						nodx.Td(
							component.TextareaControl(component.TextareaControlParams{
								Children: []nodx.Node{
									alpine.XRef("reqHeadersTextarea"),
									nodx.Text(exec.ReqHeaders.String),
								},
							}),
						),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Body")),
						nodx.Td(
							component.TextareaControl(component.TextareaControlParams{
								Children: []nodx.Node{
									alpine.XRef("reqBodyTextarea"),
									nodx.Text(exec.ReqBody.String),
								},
							}),
						),
					),
				),

				nodx.Table(
					nodx.Class("table [&_th]:text-nowrap"),
					nodx.Tr(
						nodx.Td(
							nodx.Colspan("100%"),
							component.H3Text("Response"),
						),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Status")),
						nodx.Td(component.SpanText(
							fmt.Sprintf("%d", exec.ResStatus.Int16),
						)),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Duration")),
						nodx.Td(component.SpanText(duration.String())),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Headers")),
						nodx.Td(
							component.TextareaControl(component.TextareaControlParams{
								Children: []nodx.Node{
									alpine.XRef("resHeadersTextarea"),
									nodx.Text(exec.ResHeaders.String),
								},
							}),
						),
					),
					nodx.Tr(
						nodx.Th(component.SpanText("Body")),
						nodx.Td(
							component.TextareaControl(component.TextareaControlParams{
								Children: []nodx.Node{
									alpine.XRef("resBodyTextarea"),
									nodx.Text(exec.ResBody.String),
								},
							}),
						),
					),
				),
			),
		},
	})

	return nodx.Div(
		nodx.Class("inline-block tooltip tooltip-right"),
		nodx.Data("tip", "More details"),
		mo.HTML,
		nodx.Button(
			nodx.Class("btn btn-error btn-square btn-sm btn-ghost"),
			lucide.Eye(),
			mo.OpenerAttr,
		),
	)
}

func webhookExecutionsButton(webhookID uuid.UUID) nodx.Node {
	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Webhook executions",
		Content: []nodx.Node{
			nodx.Table(
				nodx.Class("table"),
				nodx.Thead(
					nodx.Tr(
						nodx.Th(nodx.Class("w-1")),
						nodx.Th(component.SpanText("Status")),
						nodx.Th(component.SpanText("Method")),
						nodx.Th(component.SpanText("Duration")),
						nodx.Th(component.SpanText("Date")),
					),
				),
				nodx.Tbody(
					htmx.HxGet(
						"/dashboard/webhooks/"+webhookID.String()+"/executions?page=1",
					),
					htmx.HxIndicator("#webhook-executions-loading"),
					htmx.HxTrigger("intersect once"),
				),
			),
			nodx.Div(
				nodx.Class("flex justify-center pt-2"),
				component.HxLoadingMd("webhook-executions-loading"),
			),
		},
	})

	return nodx.Div(
		mo.HTML,
		component.OptionsDropdownButton(
			mo.OpenerAttr,
			lucide.List(),
			component.SpanText("Show executions"),
		),
	)
}
