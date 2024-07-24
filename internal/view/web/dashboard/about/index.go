package about

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	return echoutil.RenderGomponent(c, http.StatusOK, indexPage())
}

func indexPage() gomponents.Node {
	content := []gomponents.Node{
		component.H1Text("About PG Back Web"),
		component.H2Text(config.Version),

		html.Div(
			html.Class("grid grid-cols-2 gap-4 mt-4"),

			component.CardBox(component.CardBoxParams{
				Children: []gomponents.Node{
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
				Children: []gomponents.Node{
					html.Table(
						html.Class("table"),
						html.Tr(
							html.Th(component.SpanText("License")),
							html.Td(
								html.A(
									html.Class("link"),
									html.Href("https://github.com/eduardolat/pgbackweb/blob/main/LICENSE"),
									html.Target("_blank"),
									component.SpanText("MIT"),
								),
							),
						),
						html.Tr(
							html.Th(component.SpanText("About the author")),
							html.Td(
								html.A(
									html.Class("link"),
									html.Href("https://eduardo.lat"),
									html.Target("_blank"),
									component.SpanText("https://eduardo.lat"),
								),
							),
						),
						html.Tr(
							html.Th(component.SpanText("Repository")),
							html.Td(
								html.A(
									html.Class("link"),
									html.Href("https://github.com/eduardolat/pgbackweb"),
									html.Target("_blank"),
									component.SpanText("https://github.com/eduardolat/pgbackweb"),
								),
							),
						),
					),
				},
			}),
		),
	}

	return layout.Dashboard(layout.DashboardParams{
		Title: "About",
		Body:  content,
	})
}
