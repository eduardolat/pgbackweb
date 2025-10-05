package view

import (
	"io/fs"
	"time"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/view/api"
	"github.com/eduardolat/pgbackweb/internal/view/middleware"
	"github.com/eduardolat/pgbackweb/internal/view/static"
	"github.com/eduardolat/pgbackweb/internal/view/web"
	"github.com/labstack/echo/v4"
)

func MountRouter(app *echo.Echo, servs *service.Service) {
	mids := middleware.New(servs)

	// Create the base group with the path prefix (if any)
	baseGroup := app.Group(pathutil.GetPathPrefix())

	browserCache := mids.NewBrowserCacheMiddleware(
		middleware.BrowserCacheMiddlewareConfig{
			CacheDuration: time.Hour * 24 * 30,
			ExcludedFiles: []string{"/robots.txt"},
		},
	)

	// Mount static files
	staticFS, err := fs.Sub(static.StaticFs, ".")
	if err != nil {
		logger.FatalError("failed to create static filesystem", logger.KV{"error": err})
	}

	staticGroup := baseGroup.Group("", browserCache)
	staticGroup.StaticFS("/", staticFS)

	apiGroup := baseGroup.Group("/api")
	api.MountRouter(apiGroup, mids, servs)

	webGroup := baseGroup.Group("", mids.InjectReqctx)
	web.MountRouter(webGroup, mids, servs)
}
