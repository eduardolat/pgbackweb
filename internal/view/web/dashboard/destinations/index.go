package destinations

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
			nodx.Class("flex justify-between items-start space-x-2"),
			nodx.Div(
				component.H1Text("S3 Destinations"),
				component.PText(`
					Here you can manage your S3 destinations. You can skip creating a S3
					destination if you want to use the local storage for your backups.
				`),
			),
			nodx.Div(
				nodx.Class("flex-none"),
				createDestinationButton(),
			),
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
								nodx.Th(component.SpanText("Bucket name")),
								nodx.Th(component.SpanText("Endpoint")),
								nodx.Th(component.SpanText("Region")),
								nodx.Th(component.SpanText("Access key")),
								nodx.Th(component.SpanText("Secret key")),
								nodx.Th(component.SpanText("Created at")),
							),
						),
						nodx.Tbody(
							component.SkeletonTr(8),
							htmx.HxGet(pathutil.BuildPath("/dashboard/destinations/list?page=1")),
							htmx.HxTrigger("load"),
						),
					),
				),
			},
		}),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: "S3 Destinations",
		Body:  content,
	})
}
