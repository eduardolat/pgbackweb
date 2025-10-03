package destinations

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/destinations"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) listDestinationsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData struct {
		Page int `query:"page" validate:"required,min=1"`
	}
	if err := c.Bind(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	pagination, destinations, err := h.servs.DestinationsService.PaginateDestinations(
		ctx, destinations.PaginateDestinationsParams{
			Page:  formData.Page,
			Limit: 20,
		},
	)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return echoutil.RenderNodx(
		c, http.StatusOK, listDestinations(pagination, destinations),
	)
}

func listDestinations(
	pagination paginateutil.PaginateResponse,
	destinations []dbgen.DestinationsServicePaginateDestinationsRow,
) nodx.Node {
	if len(destinations) < 1 {
		return component.EmptyResultsTr(component.EmptyResultsParams{
			Title:    "No destinations found",
			Subtitle: "Wait for the first destination to appear here",
		})
	}

	trs := []nodx.Node{}
	for _, destination := range destinations {
		trs = append(trs, nodx.Tr(
			nodx.Td(component.OptionsDropdown(
				component.OptionsDropdownA(
					nodx.Href(
						fmt.Sprintf("/dashboard/executions?destination=%s", destination.ID),
					),
					nodx.Target("_blank"),
					lucide.List(),
					component.SpanText("Show executions"),
				),
				editDestinationButton(destination),
				component.OptionsDropdownButton(
					htmx.HxPost(pathutil.BuildPath("/dashboard/destinations/"+destination.ID.String()+"/test")),
					htmx.HxDisabledELT("this"),
					lucide.PlugZap(),
					component.SpanText("Test connection"),
				),
				deleteDestinationButton(destination.ID),
			)),
			nodx.Td(
				nodx.Div(
					nodx.Class("flex items-center space-x-2"),
					component.HealthStatusPing(
						destination.TestOk, destination.TestError, destination.LastTestAt,
					),
					component.SpanText(destination.Name),
				),
			),
			nodx.Td(
				nodx.Div(
					nodx.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.BucketName),
					component.SpanText(destination.BucketName),
				),
			),
			nodx.Td(
				nodx.Div(
					nodx.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.Endpoint),
					component.SpanText(destination.Endpoint),
				),
			),
			nodx.Td(
				nodx.Div(
					nodx.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.Region),
					component.SpanText(destination.Region),
				),
			),
			nodx.Td(
				nodx.Div(
					nodx.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.DecryptedAccessKey),
					component.SpanText("**********"),
				),
			),
			nodx.Td(
				nodx.Div(
					nodx.Class("flex items-center space-x-1"),
					component.CopyButtonSm(destination.DecryptedSecretKey),
					component.SpanText("**********"),
				),
			),
			nodx.Td(component.SpanText(
				destination.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, nodx.Tr(
			htmx.HxGet(fmt.Sprintf(
				"/dashboard/destinations/list?page=%d", pagination.NextPage,
			)),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
		))
	}

	return component.RenderableGroup(trs)
}
