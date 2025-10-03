package databases

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/databases"
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

func (h *handlers) listDatabasesHandler(c echo.Context) error {
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

	pagination, databases, err := h.servs.DatabasesService.PaginateDatabases(
		ctx, databases.PaginateDatabasesParams{
			Page:  formData.Page,
			Limit: 20,
		},
	)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return echoutil.RenderNodx(
		c, http.StatusOK, listDatabases(pagination, databases),
	)
}

func listDatabases(
	pagination paginateutil.PaginateResponse,
	databases []dbgen.DatabasesServicePaginateDatabasesRow,
) nodx.Node {
	if len(databases) < 1 {
		return component.EmptyResultsTr(component.EmptyResultsParams{
			Title:    "No databases found",
			Subtitle: "Wait for the first database to appear here",
		})
	}

	trs := []nodx.Node{}
	for _, database := range databases {
		trs = append(trs, nodx.Tr(
			nodx.Td(component.OptionsDropdown(
				nodx.Div(
					nodx.Class("flex flex-col space-y-1"),
					component.OptionsDropdownA(
						nodx.Href(
							fmt.Sprintf("/dashboard/executions?database=%s", database.ID),
						),
						nodx.Target("_blank"),
						lucide.List(),
						component.SpanText("Show executions"),
					),
					editDatabaseButton(database),
					component.OptionsDropdownButton(
						htmx.HxPost(pathutil.BuildPath("/dashboard/databases/"+database.ID.String()+"/test")),
						htmx.HxDisabledELT("this"),
						lucide.DatabaseZap(),
						component.SpanText("Test connection"),
					),
					deleteDatabaseButton(database.ID),
				),
			)),
			nodx.Td(
				nodx.Div(
					nodx.Class("flex items-center space-x-2"),
					component.HealthStatusPing(
						database.TestOk, database.TestError, database.LastTestAt,
					),
					component.SpanText(database.Name),
				),
			),
			nodx.Td(component.SpanText("PostgreSQL "+database.PgVersion)),
			nodx.Td(
				nodx.Class("space-x-1"),
				component.CopyButtonSm(database.DecryptedConnectionString),
				component.SpanText("****************"),
			),
			nodx.Td(component.SpanText(
				database.CreatedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
			)),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, nodx.Tr(
			htmx.HxGet(fmt.Sprintf(
				"/dashboard/databases/list?page=%d", pagination.NextPage,
			)),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
		))
	}

	return component.RenderableGroup(trs)
}
