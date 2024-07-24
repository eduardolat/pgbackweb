package auth

import (
	"time"

	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	servs *service.Service
}

func MountRouter(
	parent *echo.Group, mids *middleware.Middleware, servs *service.Service,
) {
	h := handlers{servs: servs}

	requireAuth := parent.Group("", mids.RequireAuth)
	requireNoAuth := parent.Group("", mids.RequireNoAuth)

	requireNoAuth.GET("/create-first-user", h.createFirstUserPageHandler)
	requireNoAuth.POST("/create-first-user", h.createFirstUserHandler)

	requireNoAuth.GET("/login", h.loginPageHandler)
	requireNoAuth.POST("/login", h.loginHandler, mids.RateLimit(middleware.RateLimitConfig{
		Limit:  5,
		Period: 10 * time.Second,
	}))

	requireAuth.POST("/logout", h.logoutHandler)
	requireAuth.POST("/logout-all", h.logoutAllSessionsHandler)
}
