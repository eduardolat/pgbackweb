package auth

import (
	"time"

	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/eduardolat/pgbackweb/internal/view/web/oidc"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	servs *service.Service
}

// MountRouter registers authentication-related HTTP routes on the provided Echo group, applying appropriate middleware for authenticated and unauthenticated access. It also mounts OpenID Connect (OIDC) routes on the same group.
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

	// Mount OIDC routes
	oidc.MountRouter(parent, mids, servs)
}
