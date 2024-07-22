package web

import (
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/web/auth"
	"github.com/labstack/echo/v4"
)

func MountRouter(parent *echo.Group, servs *service.Service) {
	auth.MountRouter(parent.Group("/auth"), servs)
}
