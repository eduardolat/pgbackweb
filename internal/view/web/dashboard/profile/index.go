package profile

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

	sessions, err := h.servs.AuthService.GetUserSessions(ctx, reqCtx.User.ID)
	if err != nil {
		logger.Error("failed to get user sessions", logger.KV{"err": err})
		return c.String(http.StatusInternalServerError, "failed to get user sessions")
	}

	return echoutil.RenderGomponent(
		c, http.StatusOK, indexPage(reqCtx.User, sessions),
	)
}

func indexPage(user dbgen.User, sessions []dbgen.Session) gomponents.Node {
	content := []gomponents.Node{
		component.H1Text("Profile"),

		html.Div(
			html.Class("mt-4 grid grid-cols-2 gap-4"),
			html.Div(updateUserForm(user)),
			html.Div(closeAllSessionsForm(sessions)),
		),
	}

	return layout.Dashboard(layout.DashboardParams{
		Title: "Profile",
		Body:  content,
	})
}
