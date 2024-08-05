package restorations

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/restorations"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type listResQueryData struct {
	Execution uuid.UUID `query:"execution" validate:"omitempty,uuid"`
	Database  uuid.UUID `query:"database" validate:"omitempty,uuid"`
	Page      int       `query:"page" validate:"required,min=1"`
}

func (h *handlers) listRestorationsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var queryData listResQueryData
	if err := c.Bind(&queryData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&queryData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	pagination, restorations, err := h.servs.RestorationsService.PaginateRestorations(
		ctx, restorations.PaginateRestorationsParams{
			ExecutionFilter: uuid.NullUUID{
				UUID: queryData.Execution, Valid: queryData.Execution != uuid.Nil,
			},
			DatabaseFilter: uuid.NullUUID{
				UUID: queryData.Database, Valid: queryData.Database != uuid.Nil,
			},
			Page:  queryData.Page,
			Limit: 20,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, listRestorations(queryData, pagination, restorations),
	)
}

func listRestorations(
	queryData listResQueryData,
	pagination paginateutil.PaginateResponse,
	restorations []dbgen.RestorationsServicePaginateRestorationsRow,
) gomponents.Node {
	trs := []gomponents.Node{}
	for _, restoration := range restorations {
		trs = append(trs, html.Tr(
			html.Td(
				html.Class("w-[50px]"),
				html.Div(
					html.Class("flex justify-start space-x-1"),
					showRestorationButton(restoration),
				),
			),
			html.Td(component.StatusBadge(restoration.Status)),
			html.Td(component.SpanText(restoration.BackupName)),
			html.Td(component.SpanText(func() string {
				if restoration.DatabaseName.Valid {
					return restoration.DatabaseName.String
				}
				return "Other database"
			}())),
			html.Td(component.SpanText(restoration.ExecutionID.String())),
			html.Td(component.SpanText(
				restoration.StartedAt.Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
			html.Td(
				gomponents.If(
					restoration.FinishedAt.Valid,
					component.SpanText(
						restoration.FinishedAt.Time.Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
					),
				),
			),
			html.Td(
				gomponents.If(
					restoration.FinishedAt.Valid,
					component.SpanText(
						restoration.FinishedAt.Time.Sub(restoration.StartedAt).String(),
					),
				),
			),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, html.Tr(
			htmx.HxGet(func() string {
				url := "/dashboard/restorations/list"
				url = strutil.AddQueryParamToUrl(url, "page", fmt.Sprintf("%d", pagination.NextPage))
				if queryData.Execution != uuid.Nil {
					url = strutil.AddQueryParamToUrl(url, "execution", queryData.Execution.String())
				}
				if queryData.Database != uuid.Nil {
					url = strutil.AddQueryParamToUrl(url, "database", queryData.Database.String())
				}
				return url
			}()),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
			htmx.HxIndicator("#list-restorations-loading"),
		))
	}

	return component.RenderableGroup(trs)
}
