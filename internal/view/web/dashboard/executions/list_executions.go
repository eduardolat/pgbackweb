package executions

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/executions"
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

type listExecsQueryData struct {
	Database    uuid.UUID `query:"database" validate:"omitempty,uuid"`
	Destination uuid.UUID `query:"destination" validate:"omitempty,uuid"`
	Backup      uuid.UUID `query:"backup" validate:"omitempty,uuid"`
	Page        int       `query:"page" validate:"required,min=1"`
}

func (h *handlers) listExecutionsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var queryData listExecsQueryData
	if err := c.Bind(&queryData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}
	if err := validate.Struct(&queryData); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	pagination, executions, err := h.servs.ExecutionsService.PaginateExecutions(
		ctx, executions.PaginateExecutionsParams{
			DatabaseFilter: uuid.NullUUID{
				UUID: queryData.Database, Valid: queryData.Database != uuid.Nil,
			},
			DestinationFilter: uuid.NullUUID{
				UUID: queryData.Destination, Valid: queryData.Destination != uuid.Nil,
			},
			BackupFilter: uuid.NullUUID{
				UUID: queryData.Backup, Valid: queryData.Backup != uuid.Nil,
			},
			Page:  queryData.Page,
			Limit: 20,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, listExecutions(queryData, pagination, executions),
	)
}

func listExecutions(
	queryData listExecsQueryData,
	pagination paginateutil.PaginateResponse,
	executions []dbgen.ExecutionsServicePaginateExecutionsRow,
) gomponents.Node {
	trs := []gomponents.Node{}
	for _, execution := range executions {
		trs = append(trs, html.Tr(
			html.Td(
				html.Class("w-[50px]"),
				html.Div(
					html.Class("flex justify-start space-x-1"),
					showExecutionButton(execution),
					restoreExecutionButton(execution),
				),
			),
			html.Td(component.StatusBadge(execution.Status)),
			html.Td(component.SpanText(execution.BackupName)),
			html.Td(component.SpanText(execution.DatabaseName)),
			html.Td(component.PrettyDestinationName(
				execution.BackupIsLocal, execution.DestinationName,
			)),
			html.Td(component.SpanText(
				execution.StartedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
			html.Td(
				gomponents.If(
					execution.FinishedAt.Valid,
					component.SpanText(
						execution.FinishedAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
					),
				),
			),
			html.Td(
				gomponents.If(
					execution.FinishedAt.Valid,
					component.SpanText(
						execution.FinishedAt.Time.Sub(execution.StartedAt).String(),
					),
				),
			),
			html.Td(
				gomponents.If(
					execution.FileSize.Valid,
					component.PrettyFileSize(execution.FileSize),
				),
			),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, html.Tr(
			htmx.HxGet(func() string {
				url := "/dashboard/executions/list"
				url = strutil.AddQueryParamToUrl(url, "page", fmt.Sprintf("%d", pagination.NextPage))
				if queryData.Database != uuid.Nil {
					url = strutil.AddQueryParamToUrl(url, "database", queryData.Database.String())
				}
				if queryData.Destination != uuid.Nil {
					url = strutil.AddQueryParamToUrl(url, "destination", queryData.Destination.String())
				}
				if queryData.Backup != uuid.Nil {
					url = strutil.AddQueryParamToUrl(url, "backup", queryData.Backup.String())
				}
				return url
			}()),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
			htmx.HxIndicator("#list-executions-loading"),
		))
	}

	return component.RenderableGroup(trs)
}
