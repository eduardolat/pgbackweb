package restorations

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

type resQueryData struct {
	Execution uuid.UUID `query:"execution" validate:"omitempty,uuid"`
	Database  uuid.UUID `query:"database" validate:"omitempty,uuid"`
}

func (h *handlers) indexPageHandler(c echo.Context) error {
	reqCtx := reqctx.GetCtx(c)

	var queryData resQueryData
	if err := c.Bind(&queryData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(&queryData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return echoutil.RenderGomponent(c, http.StatusOK, indexPage(reqCtx, queryData))
}

func indexPage(reqCtx reqctx.Ctx, queryData resQueryData) gomponents.Node {
	content := []gomponents.Node{
		component.H1Text("Restorations"),
		component.CardBox(component.CardBoxParams{
			Class: "mt-4",
			Children: []gomponents.Node{
				html.Div(
					html.Class("overflow-x-auto"),
					html.Table(
						html.Class("table text-nowrap"),
						html.THead(
							html.Tr(
								html.Th(html.Class("w-1")),
								html.Th(component.SpanText("Status")),
								html.Th(component.SpanText("Backup")),
								html.Th(component.SpanText("Database")),
								html.Th(component.SpanText("Execution")),
								html.Th(component.SpanText("Started at")),
								html.Th(component.SpanText("Finished at")),
								html.Th(component.SpanText("Duration")),
							),
						),
						html.TBody(
							component.SkeletonTr(8),
							htmx.HxGet(func() string {
								url := "/dashboard/restorations/list?page=1"
								if queryData.Execution != uuid.Nil {
									url = strutil.AddQueryParamToUrl(url, "execution", queryData.Execution.String())
								}
								if queryData.Database != uuid.Nil {
									url = strutil.AddQueryParamToUrl(url, "database", queryData.Database.String())
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
		Title: "Restorations",
		Body:  content,
	})
}
