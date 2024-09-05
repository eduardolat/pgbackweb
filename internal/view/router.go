package view

import (
	"time"

	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/api"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/eduardolat/pgbackweb/internal/view/static"
	"github.com/eduardolat/pgbackweb/internal/view/web"
	"github.com/labstack/echo/v4"
)

func MountRouter(app *echo.Echo, servs *service.Service) {
	mids := middleware.New(servs)

	browserCache := mids.NewBrowserCacheMiddleware(
		middleware.BrowserCacheMiddlewareConfig{
			CacheDuration: time.Hour * 24 * 30,
			ExcludedFiles: []string{"/robots.txt"},
		},
	)
	app.Group("", browserCache).StaticFS("", static.StaticFs)

	apiGroup := app.Group("/api")
	api.MountRouter(apiGroup, mids, servs)

	webGroup := app.Group("", mids.InjectReqctx)
	web.MountRouter(webGroup, mids, servs)
}
