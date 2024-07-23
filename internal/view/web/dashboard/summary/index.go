package summary

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	return echoutil.RenderGomponent(c, http.StatusOK, indexPage())
}

func indexPage() gomponents.Node {
	content := []gomponents.Node{
		component.H1Text("Summary"),
	}

	return layout.Dashboard(layout.DashboardParams{
		Title: "Summary",
		Body:  content,
	})
}
