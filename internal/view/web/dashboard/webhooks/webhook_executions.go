package webhooks

import (
	"fmt"
	"net/http"
	"time"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/alpine"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) paginateWebhookExecutionsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	webhookID, err := uuid.Parse(c.Param("webhookID"))
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	var queryData struct {
		Page int `query:"page" validate:"required,min=1"`
	}
	if err := c.Bind(&queryData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&queryData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	pagination, execs, err := h.servs.WebhooksService.PaginateWebhookExecutions(
		ctx, webhooks.PaginateWebhookExecutionsParams{
			WebhookID: webhookID,
			Page:      queryData.Page,
			Limit:     20,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, webhookExecutionsList(webhookID, pagination, execs),
	)
}

func webhookExecutionsList(
	webhookID uuid.UUID,
	pagination paginateutil.PaginateResponse,
	execs []dbgen.WebhookExecution,
) gomponents.Node {
	if len(execs) == 0 {
		return component.EmptyResultsTr(component.EmptyResultsParams{
			Title:    "No executions found",
			Subtitle: "Wait for the first execution to appear here",
		})
	}

	trs := []gomponents.Node{}
	for _, exec := range execs {
		durationMillis := exec.ResDuration.Int32
		duration := time.Duration(durationMillis) * time.Millisecond

		trs = append(trs, html.Tr(
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-2"),
					webhookExecutionDetailsButton(exec, duration),
					component.SpanText(fmt.Sprintf("%d", exec.ResStatus.Int16)),
				),
			),
			html.Td(component.SpanText(exec.ReqMethod.String)),
			html.Td(component.SpanText(duration.String())),
			html.Td(component.SpanText(
				exec.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, html.Tr(
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
) gomponents.Node {
	mo := component.Modal(component.ModalParams{
		Title: "Webhook execution details",
		Content: []gomponents.Node{
			html.Div(
				html.Class("space-y-4"),

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

				html.Table(
					html.Class("table [&_th]:text-nowrap"),
					html.Tr(
						html.Td(
							html.ColSpan("100%"),
							component.H3Text("General"),
						),
					),
					html.Tr(
						html.Th(component.SpanText("ID")),
						html.Td(component.SpanText(exec.ID.String())),
					),
					html.Tr(
						html.Th(component.SpanText("Date")),
						html.Td(component.SpanText(
							exec.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
						)),
					),
				),

				html.Table(
					html.Class("table [&_th]:text-nowrap"),
					html.Tr(
						html.Td(
							html.ColSpan("100%"),
							component.H3Text("Request"),
						),
					),
					html.Tr(
						html.Th(component.SpanText("Method")),
						html.Td(component.SpanText(exec.ReqMethod.String)),
					),
					html.Tr(
						html.Th(component.SpanText("Headers")),
						html.Td(
							component.TextareaControl(component.TextareaControlParams{
								Children: []gomponents.Node{
									alpine.XRef("reqHeadersTextarea"),
									gomponents.Text(exec.ReqHeaders.String),
								},
							}),
						),
					),
					html.Tr(
						html.Th(component.SpanText("Body")),
						html.Td(
							component.TextareaControl(component.TextareaControlParams{
								Children: []gomponents.Node{
									alpine.XRef("reqBodyTextarea"),
									gomponents.Text(exec.ReqBody.String),
								},
							}),
						),
					),
				),

				html.Table(
					html.Class("table [&_th]:text-nowrap"),
					html.Tr(
						html.Td(
							html.ColSpan("100%"),
							component.H3Text("Response"),
						),
					),
					html.Tr(
						html.Th(component.SpanText("Status")),
						html.Td(component.SpanText(
							fmt.Sprintf("%d", exec.ResStatus.Int16),
						)),
					),
					html.Tr(
						html.Th(component.SpanText("Duration")),
						html.Td(component.SpanText(duration.String())),
					),
					html.Tr(
						html.Th(component.SpanText("Headers")),
						html.Td(
							component.TextareaControl(component.TextareaControlParams{
								Children: []gomponents.Node{
									alpine.XRef("resHeadersTextarea"),
									gomponents.Text(exec.ResHeaders.String),
								},
							}),
						),
					),
					html.Tr(
						html.Th(component.SpanText("Body")),
						html.Td(
							component.TextareaControl(component.TextareaControlParams{
								Children: []gomponents.Node{
									alpine.XRef("resBodyTextarea"),
									gomponents.Text(exec.ResBody.String),
								},
							}),
						),
					),
				),
			),
		},
	})

	return html.Div(
		html.Class("inline-block tooltip tooltip-right"),
		html.Data("tip", "Show webhook execution details"),
		mo.HTML,
		html.Button(
			html.Class("btn btn-error btn-square btn-sm btn-ghost"),
			lucide.Eye(),
			mo.OpenerAttr,
		),
	)
}

func webhookExecutionsButton(webhookID uuid.UUID) gomponents.Node {
	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: "Webhook executions",
		Content: []gomponents.Node{
			html.Table(
				html.Class("table"),
				html.THead(
					html.Tr(
						html.Th(
							html.Div(
								html.Class("ml-10"),
								component.SpanText("Status"),
							),
						),
						html.Th(component.SpanText("Method")),
						html.Th(component.SpanText("Duration")),
						html.Th(component.SpanText("Date")),
					),
				),
				html.TBody(
					htmx.HxGet(
						"/dashboard/webhooks/"+webhookID.String()+"/executions?page=1",
					),
					htmx.HxIndicator("#webhook-executions-loading"),
					htmx.HxTrigger("intersect once"),
				),
			),
			html.Div(
				html.Class("flex justify-center pt-2"),
				component.HxLoadingMd("webhook-executions-loading"),
			),
		},
	})

	return html.Div(
		html.Class("inline-block tooltip tooltip-right"),
		html.Data("tip", "Show webhook executions"),
		mo.HTML,
		html.Button(
			html.Class("btn btn-error btn-square btn-sm btn-ghost"),
			lucide.List(),
			mo.OpenerAttr,
		),
	)
}
