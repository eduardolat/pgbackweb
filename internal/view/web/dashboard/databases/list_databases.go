package databases

import (
	"fmt"
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/databases"
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

func (h *handlers) listDatabasesHandler(c echo.Context) error {
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

	pagination, databases, err := h.servs.DatabasesService.PaginateDatabases(
		ctx, databases.PaginateDatabasesParams{
			Page:  formData.Page,
			Limit: 20,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, listDatabases(pagination, databases),
	)
}

func listDatabases(
	pagination paginateutil.PaginateResponse,
	databases []dbgen.DatabasesServicePaginateDatabasesRow,
) gomponents.Node {
	if len(databases) < 1 {
		return component.EmptyResultsTr(component.EmptyResultsParams{
			Title:    "No databases found",
			Subtitle: "Wait for the first database to appear here",
		})
	}

	trs := []gomponents.Node{}
	for _, database := range databases {
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
								fmt.Sprintf("/dashboard/executions?database=%s", database.ID),
							),
							html.Target("_blank"),
							lucide.List(),
						),
					),
					editDatabaseButton(database),
					html.Div(
						html.Class("inline-block tooltip tooltip-right"),
						html.Data("tip", "Test connection"),
						html.Button(
							htmx.HxPost("/dashboard/databases/"+database.ID.String()+"/test"),
							htmx.HxDisabledELT("this"),
							html.Class("btn btn-neutral btn-square btn-ghost btn-sm"),
							lucide.DatabaseZap(),
						),
					),
					deleteDatabaseButton(database.ID),
				),
			),
			html.Td(
				html.Div(
					html.Class("flex items-center space-x-2"),
					component.HealthStatusPing(
						database.TestOk, database.TestError, database.LastTestAt,
					),
					component.SpanText(database.Name),
				),
			),
			html.Td(component.SpanText("PostgreSQL "+database.PgVersion)),
			html.Td(
				html.Class("space-x-1"),
				component.CopyButtonSm(database.DecryptedConnectionString),
				component.SpanText("****************"),
			),
			html.Td(component.SpanText(
				database.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, html.Tr(
			htmx.HxGet(fmt.Sprintf(
				"/dashboard/databases/list?page=%d", pagination.NextPage,
			)),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
		))
	}

	return component.RenderableGroup(trs)
}
