package auth

import (
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmxs"
	"github.com/labstack/echo/v4"
)

func (h *handlers) logoutHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

	if err := h.servs.AuthService.DeleteSession(ctx, reqCtx.SessionID); err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	h.servs.AuthService.ClearSessionCookie(c)
	return htmxs.RespondRedirect(c, "/auth/login")
}

func (h *handlers) logoutAllSessionsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

	err := h.servs.AuthService.DeleteAllUserSessions(ctx, reqCtx.User.ID)
	if err != nil {
		return htmxs.RespondToastError(c, err.Error())
	}

	h.servs.AuthService.ClearSessionCookie(c)
	return htmxs.RespondRedirect(c, "/auth/login")
}
