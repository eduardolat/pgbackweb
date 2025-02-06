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
	nodx "github.com/nodxdev/nodxgo"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

	sessions, err := h.servs.AuthService.GetUserSessions(ctx, reqCtx.User.ID)
	if err != nil {
		logger.Error("failed to get user sessions", logger.KV{"err": err})
		return c.String(http.StatusInternalServerError, "failed to get user sessions")
	}

	return echoutil.RenderNodx(
		c, http.StatusOK, indexPage(reqCtx, sessions),
	)
}

func indexPage(reqCtx reqctx.Ctx, sessions []dbgen.Session) nodx.Node {
	content := []nodx.Node{
		component.H1Text("Profile"),

		nodx.Div(
			nodx.Class("mt-4 grid grid-cols-2 gap-4"),
			nodx.Div(updateUserForm(reqCtx.User)),
			nodx.Div(closeAllSessionsForm(sessions)),
		),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: "Profile",
		Body:  content,
	})
}
