package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) loginPageHandler(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}
