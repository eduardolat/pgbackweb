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
					editDestinationButton(destination),
					html.Form(
						html.Class("inline-block tooltip tooltip-right"),
						html.Data("tip", "Test connection"),
						htmx.HxPost("/dashboard/destinations/test"),
						html.Input(
							html.Type("hidden"),
							html.Name("name"),
							html.Value(destination.Name),
						),
						html.Input(
							html.Type("hidden"),
							html.Name("bucket_name"),
							html.Value(destination.BucketName),
						),
						html.Input(
							html.Type("hidden"),
							html.Name("endpoint"),
							html.Value(destination.Endpoint),
						),
						html.Input(
							html.Type("hidden"),
							html.Name("region"),
							html.Value(destination.Region),
						),
						html.Input(
							html.Type("hidden"),
							html.Name("access_key"),
							html.Value(destination.DecryptedAccessKey),
						),
						html.Input(
							html.Type("hidden"),
							html.Name("secret_key"),
							html.Value(destination.DecryptedSecretKey),
						),
						html.Button(
							html.Class("btn btn-neutral btn-square btn-ghost btn-sm"),
							lucide.PlugZap(),
						),
					),
					deleteDestinationButton(destination.ID),
				),
			),
			html.Td(component.SpanText(destination.Name)),
			html.Td(component.SpanText(destination.BucketName)),
			html.Td(component.SpanText(destination.Endpoint)),
			html.Td(component.SpanText(destination.Region)),
			html.Td(
				html.Class("space-x-1"),
				component.CopyButtonSm(destination.DecryptedAccessKey),
				component.SpanText("**********"),
			),
			html.Td(
				html.Class("space-x-1"),
				component.CopyButtonSm(destination.DecryptedSecretKey),
				component.SpanText("**********"),
			),
			html.Td(component.SpanText(
				destination.CreatedAt.Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
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
