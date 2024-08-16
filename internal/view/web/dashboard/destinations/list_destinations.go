package destinations

import (
	"fmt"
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/destinations"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) listDestinationsHandler(c echo.Context) error {
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

	pagination, destinations, err := h.servs.DestinationsService.PaginateDestinations(
		ctx, destinations.PaginateDestinationsParams{
			Page:  formData.Page,
			Limit: 20,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, listDestinations(pagination, destinations),
	)
}

func listDestinations(
	pagination paginateutil.PaginateResponse,
	destinations []dbgen.DestinationsServicePaginateDestinationsRow,
) gomponents.Node {
	trs := []gomponents.Node{}
	for _, destination := range destinations {
		trs = append(trs, html.Tr(
			html.Td(
				html.Class("w-[40px]"),
				html.Div(
					html.Class("flex justify-start space-x-1"),
					html.Div(
						html.Class("inline-block tooltip tooltip-right"),
						html.Data("tip", "Show executions"),
						html.A(
							html.Class("btn btn-sm btn-ghost btn-square"),
							html.Href(
								fmt.Sprintf("/dashboard/executions?destination=%s", destination.ID),
							),
							html.Target("_blank"),
							lucide.List(),
						),
					),
					editDestinationButton(destination),
					html.Div(
						html.Class("inline-block tooltip tooltip-right"),
						html.Data("tip", "Test connection"),
						html.Button(
							htmx.HxPost("/dashboard/destinations/"+destination.ID.String()+"/test"),
							htmx.HxDisabledELT("this"),
							html.Class("btn btn-neutral btn-square btn-ghost btn-sm"),
							lucide.PlugZap(),
						),
					),
					deleteDestinationButton(destination.ID),
				),
			),
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-2"),
					component.HealthStatusPing(
						destination.TestOk, destination.TestError, destination.LastTestAt,
					),
					component.SpanText(destination.Name),
				),
			),
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.BucketName),
					component.SpanText(destination.BucketName),
				),
			),
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.Endpoint),
					component.SpanText(destination.Endpoint),
				),
			),
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.Region),
					component.SpanText(destination.Region),
				),
			),
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.DecryptedAccessKey),
					component.SpanText("**********"),
				),
			),
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.DecryptedSecretKey),
					component.SpanText("**********"),
				),
			),
			html.Td(component.SpanText(
				destination.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, html.Tr(
			htmx.HxGet(fmt.Sprintf(
				"/dashboard/destinations/list?page=%d", pagination.NextPage,
			)),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
			htmx.HxIndicator("#list-destinations-loading"),
		))
	}

	return component.RenderableGroup(trs)
}
