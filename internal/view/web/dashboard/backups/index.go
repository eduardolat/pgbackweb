package backups

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	reqCtx := reqctx.GetCtx(c)
	return echoutil.RenderNodx(c, http.StatusOK, indexPage(reqCtx))
}

func indexPage(reqCtx reqctx.Ctx) nodx.Node {
	content := []nodx.Node{
		nodx.Div(
			nodx.Class("flex justify-between items-start"),
			component.H1Text("Backup tasks"),
			createBackupButton(),
		),
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
								nodx.Th(component.SpanText("Name")),
								nodx.Th(component.SpanText("Database")),
								nodx.Th(component.SpanText("Destination")),
								nodx.Th(component.SpanText("Schedule")),
								nodx.Th(component.SpanText("Retention")),
								nodx.Th(component.SpanText("--data-only")),
								nodx.Th(component.SpanText("--schema-only")),
								nodx.Th(component.SpanText("--clean")),
								nodx.Th(component.SpanText("--if-exists")),
								nodx.Th(component.SpanText("--create")),
								nodx.Th(component.SpanText("--no-comments")),
								nodx.Th(component.SpanText("Created at")),
							),
						),
						nodx.Tbody(
							component.SkeletonTr(8),
							htmx.HxGet(pathutil.BuildPath("/dashboard/backups/list?page=1")),
							htmx.HxTrigger("load"),
						),
					),
				),
			},
		}),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: "Backup tasks",
		Body:  content,
	})
}
