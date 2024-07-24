package databases

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/databases"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexListDatabasesHandler(c echo.Context) error {
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
			Limit: 8,
		},
	)
	if err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, indexListDatabases(pagination, databases),
	)
}

func indexListDatabases(
	pagination paginateutil.PaginateResponse,
	databases []dbgen.DatabasesServicePaginateDatabasesRow,
) gomponents.Node {
	trs := []gomponents.Node{}
	for _, database := range databases {
		trs = append(trs, html.Tr(
			html.Td(),
			html.Td(component.SpanText(database.Name)),
			html.Td(component.SpanText("PostgreSQL "+database.PgVersion)),
			html.Td(
				html.Class("space-x-1"),
				component.CopyButtonSm(database.DecryptedConnectionString),
				component.SpanText("****************"),
			),
		))
	}

	if pagination.HasNextPage {
		trs = append(trs, html.Tr(
			htmx.HxGet(fmt.Sprintf(
				"/dashboard/databases/list?page=%d", pagination.NextPage,
			)),
			htmx.HxTrigger("intersect once"),
			htmx.HxSwap("afterend"),
			htmx.HxIndicator("#list-databases-loading"),
		))
	}

	return component.RenderableGroup(trs)
}
