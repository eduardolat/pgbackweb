package executions

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type execsQueryData struct {
	Database    uuid.UUID `query:"database" validate:"omitempty,uuid"`
	Destination uuid.UUID `query:"destination" validate:"omitempty,uuid"`
	Backup      uuid.UUID `query:"backup" validate:"omitempty,uuid"`
}

func (h *handlers) indexPageHandler(c echo.Context) error {
	reqCtx := reqctx.GetCtx(c)

	var queryData execsQueryData
	if err := c.Bind(&queryData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(&queryData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return echoutil.RenderGomponent(c, http.StatusOK, indexPage(reqCtx, queryData))
}

func indexPage(reqCtx reqctx.Ctx, queryData execsQueryData) gomponents.Node {
	content := []gomponents.Node{
		component.H1Text("Executions"),
		component.CardBox(component.CardBoxParams{
			Class: "mt-4",
			Children: []gomponents.Node{
				html.Div(
					html.Class("overflow-x-auto"),
					html.Table(
						html.Class("table text-nowrap"),
						html.THead(
							html.Tr(
								html.Th(component.SpanText("Actions")),
								html.Th(component.SpanText("Status")),
								html.Th(component.SpanText("Backup")),
								html.Th(component.SpanText("Database")),
								html.Th(component.SpanText("Destination")),
								html.Th(component.SpanText("Started at")),
								html.Th(component.SpanText("Finished at")),
								html.Th(component.SpanText("Duration")),
								html.Th(component.SpanText("File size")),
							),
						),
						html.TBody(
							htmx.HxGet(func() string {
								url := "/dashboard/executions/list?page=1"
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
							htmx.HxTrigger("load"),
						),
					),
				),
			},
		}),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: "Executions",
		Body:  content,
	})
}
