package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type BrowserCacheMiddlewareConfig struct {
	CacheDuration time.Duration
	ExcludedFiles []string
}

// NewBrowserCacheMiddleware creates a new CacheMiddleware with the specified
// cache duration and a list of excluded files that will bypass the cache.
func (Middleware) NewBrowserCacheMiddleware(
	config BrowserCacheMiddlewareConfig,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			for _, excluded := range config.ExcludedFiles {
				if path == excluded {
					return next(c)
				}
			}

			cacheDuration := config.CacheDuration
			c.Response().Header().Set(
				"Cache-Control",
				"public, max-age="+strconv.Itoa(int(cacheDuration.Seconds())),
			)

			return next(c)
		}
	}
}
