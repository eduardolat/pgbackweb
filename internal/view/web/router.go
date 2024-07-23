package web

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/auth"
	"github.com/eduardolat/pgbackweb/internal/view/web/dashboard"
	"github.com/labstack/echo/v4"
)

func MountRouter(
	parent *echo.Group, mids *middleware.Middleware, servs *service.Service,
) {
	// GET / -> Handle the root path redirects
	parent.GET("", func(c echo.Context) error {
		ctx := c.Request().Context()
		reqCtx := reqctx.GetCtx(c)

		if reqCtx.IsAuthed {
			return c.Redirect(http.StatusFound, "/dashboard")
		}

		usersQty, err := servs.UsersService.GetUsersQty(ctx)
		if err != nil {
			logger.Error("failed to get users qty", logger.KV{
				"ip":    c.RealIP(),
				"ua":    c.Request().UserAgent(),
				"error": err,
			})
			return c.String(http.StatusInternalServerError, "Internal server error")
		}

		if usersQty == 0 {
			return c.Redirect(http.StatusFound, "/auth/create-first-user")
		}

		return c.Redirect(http.StatusFound, "/auth/login")
	})

	authGroup := parent.Group("/auth")
	auth.MountRouter(authGroup, mids, servs)

	dashboardGroup := parent.Group("/dashboard", mids.RequireAuth)
	dashboard.MountRouter(dashboardGroup, mids, servs)
}
