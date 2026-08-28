package middleware

import (
	"github.com/labstack/echo/v4"
)

// NoIndex sets the X-Robots-Tag header on every response to prevent
// search engines from indexing the application pages.
func (m *Middleware) NoIndex(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Robots-Tag", "noindex, nofollow")
		return next(c)
	}
}
