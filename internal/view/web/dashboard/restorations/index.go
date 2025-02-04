package restorations

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
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

	return echoutil.RenderNodx(c, http.StatusOK, indexPage(reqCtx, queryData))
}

func indexPage(reqCtx reqctx.Ctx, queryData resQueryData) nodx.Node {
	content := []nodx.Node{
		component.H1Text("Restorations"),
		component.CardBox(component.CardBoxParams{
			Class: "mt-4",
			Children: []nodx.Node{
				nodx.Div(
					nodx.Class("overflow-x-auto"),
					nodx.Table(
						nodx.Class("table text-nowrap"),
						nodx.Thead(
							nodx.Tr(
								nodx.Th(nodx.Class("w-1")),
								nodx.Th(component.SpanText("Status")),
								nodx.Th(component.SpanText("Backup")),
								nodx.Th(component.SpanText("Database")),
								nodx.Th(component.SpanText("Execution")),
								nodx.Th(component.SpanText("Started at")),
								nodx.Th(component.SpanText("Finished at")),
								nodx.Th(component.SpanText("Duration")),
							),
						),
						nodx.Tbody(
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
