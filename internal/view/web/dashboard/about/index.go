package about

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	reqCtx := reqctx.GetCtx(c)
	return echoutil.RenderNodx(c, http.StatusOK, indexPage(reqCtx))
}

func indexPage(reqCtx reqctx.Ctx) nodx.Node {
	content := []nodx.Node{
		component.H1Text("About PG Back Web"),
		component.H2Text(config.Version),

		nodx.Div(
			nodx.Class("grid grid-cols-2 gap-4 mt-4"),

			component.CardBox(component.CardBoxParams{
				Children: []nodx.Node{
					component.PText(`
						PG Back Web was born in July 2024 out of a need for a simple and
						user-friendly backup solution for self-hosted PostgreSQL databases.
						After searching extensively for an easy-to-use backup tool and not
						finding one, I decided to create my own. Its mission is to provide a
						straightforward web interface that makes managing PostgreSQL backups
						effortless and efficient.
					`),
				},
			}),

			component.CardBox(component.CardBoxParams{
				Children: []nodx.Node{
					nodx.Table(
						nodx.Class("table"),
						nodx.Tr(
							nodx.Th(component.SpanText("License")),
							nodx.Td(
								nodx.A(
									nodx.Class("link"),
									nodx.Href("https://github.com/eduardolat/pgbackweb/blob/main/LICENSE"),
									nodx.Target("_blank"),
									component.SpanText("AGPL v3"),
								),
							),
						),
						nodx.Tr(
							nodx.Th(component.SpanText("About the author")),
							nodx.Td(
								nodx.A(
									nodx.Class("link"),
									nodx.Href("https://eduardo.lat"),
									nodx.Target("_blank"),
									component.SpanText("https://eduardo.lat"),
								),
							),
						),
						nodx.Tr(
							nodx.Th(component.SpanText("Repository")),
							nodx.Td(
								nodx.A(
									nodx.Class("link"),
									nodx.Href("https://github.com/eduardolat/pgbackweb"),
									nodx.Target("_blank"),
									component.SpanText("https://github.com/eduardolat/pgbackweb"),
								),
							),
						),
					),
				},
			}),
		),

		nodx.Div(
			nodx.Class("mt-4"),
			component.SupportProjectSponsors(),
		),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: "About",
		Body:  content,
	})
}
