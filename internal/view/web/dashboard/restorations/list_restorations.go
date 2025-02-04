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
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
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
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&queryData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
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
		return respondhtmx.ToastError(c, err.Error())
	}

	return echoutil.RenderNodx(
		c, http.StatusOK, listRestorations(queryData, pagination, restorations),
	)
}

func listRestorations(
	queryData listResQueryData,
	pagination paginateutil.PaginateResponse,
	restorations []dbgen.RestorationsServicePaginateRestorationsRow,
) nodx.Node {
	if len(restorations) < 1 {
		return component.EmptyResultsTr(component.EmptyResultsParams{
			Title:    "No restorations found",
			Subtitle: "Wait for the first restoration to appear here",
		})
	}

	trs := []nodx.Node{}
	for _, restoration := range restorations {
		trs = append(trs, nodx.Tr(
			nodx.Td(
				showRestorationButton(restoration),
			),
			nodx.Td(component.StatusBadge(restoration.Status)),
			nodx.Td(component.SpanText(restoration.BackupName)),
			nodx.Td(component.SpanText(func() string {
				if restoration.DatabaseName.Valid {
					return restoration.DatabaseName.String
				}
				return "Other database"
			}())),
			nodx.Td(component.SpanText(restoration.ExecutionID.String())),
			nodx.Td(component.SpanText(
				restoration.StartedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
			nodx.Td(
				nodx.If(
					restoration.FinishedAt.Valid,
					component.SpanText(
						restoration.FinishedAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
					),
				),
			),
			nodx.Td(
				nodx.If(
					restoration.FinishedAt.Valid,
					component.SpanText(
						restoration.FinishedAt.Time.Sub(restoration.StartedAt).String(),
					),
				),
			),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, nodx.Tr(
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
		))
	}

	return component.RenderableGroup(trs)
}
