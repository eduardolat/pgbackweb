package view

import (
	"time"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view/api"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/eduardolat/pgbackweb/internal/view/static"
	"github.com/eduardolat/pgbackweb/internal/view/web"
	"github.com/labstack/echo/v4"
)

func MountRouter(app *echo.Echo, servs *service.Service, env config.Env) {
	mids := middleware.New(servs)

	// Create the base group with the path prefix (if any)
	baseGroup := app.Group(env.PBW_PATH_PREFIX)

	browserCache := mids.NewBrowserCacheMiddleware(
		middleware.BrowserCacheMiddlewareConfig{
			CacheDuration: time.Hour * 24 * 30,
			ExcludedFiles: []string{"/robots.txt"},
		},
	)
	baseGroup.Group("", browserCache).StaticFS("", static.StaticFs)

	apiGroup := baseGroup.Group("/api")
	api.MountRouter(apiGroup, mids, servs)

	webGroup := baseGroup.Group("", mids.InjectReqctx)
	web.MountRouter(webGroup, mids, servs)
}
