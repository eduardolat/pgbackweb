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

	parent.GET("/create-first-user", h.createFirstUserPageHandler)
	parent.POST("/create-first-user", h.createFirstUserHandler)

	parent.GET("/login", h.loginPageHandler)
	parent.POST("/login", h.loginHandler, mids.RateLimit(middleware.RateLimitConfig{
		Limit:  5,
		Period: 10 * time.Second,
	}))
}
