package view

import (
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/api"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/eduardolat/pgbackweb/internal/view/static"
	"github.com/eduardolat/pgbackweb/internal/view/web"
	"github.com/labstack/echo/v4"
)

func MountRouter(app *echo.Echo, servs *service.Service) {
	mids := middleware.New(servs)

	app.StaticFS("", static.StaticFs)

	apiGroup := app.Group("/api")
	api.MountRouter(apiGroup, mids, servs)

	webGroup := app.Group("", mids.InjectReqctx)
	web.MountRouter(webGroup, mids, servs)
}
