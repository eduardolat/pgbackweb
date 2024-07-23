package dashboard

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/labstack/echo/v4"
)

func MountRouter(
	parent *echo.Group, mids *middleware.Middleware, servs *service.Service,
) {
	parent.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Dashboard")
	})
}
