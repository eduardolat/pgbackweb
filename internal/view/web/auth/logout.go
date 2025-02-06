package auth

import (
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/labstack/echo/v4"
)

func (h *handlers) logoutHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

	if err := h.servs.AuthService.DeleteSession(ctx, reqCtx.SessionID); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	h.servs.AuthService.ClearSessionCookie(c)
	return respondhtmx.Redirect(c, "/auth/login")
}

func (h *handlers) logoutAllSessionsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

	err := h.servs.AuthService.DeleteAllUserSessions(ctx, reqCtx.User.ID)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	h.servs.AuthService.ClearSessionCookie(c)
	return respondhtmx.Redirect(c, "/auth/login")
}
