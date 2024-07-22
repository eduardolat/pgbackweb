package api

import (
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/labstack/echo/v4"
)

func MountRouter(parent *echo.Group, servs *service.Service) {}
