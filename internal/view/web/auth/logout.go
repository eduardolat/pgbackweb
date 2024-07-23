package auth

import (
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/labstack/echo/v4"
)

func (h *handlers) logoutHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

	if err := h.servs.AuthService.DeleteSession(ctx, reqCtx.SessionID); err != nil {
		return htmx.RespondToastError(c, err.Error())
	}

	h.servs.AuthService.ClearSessionCookie(c)
	return htmx.RespondRedirect(c, "/auth/login")
}
