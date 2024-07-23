package auth

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/labstack/echo/v4"
)

const (
	sessionCookieName = "pbw_session"
)

func (s *Service) SetSessionCookie(c echo.Context, session dbgen.Session) {
	cookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    session.Token,
		MaxAge:   int(maxSessionAge.Seconds()),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	c.SetCookie(&cookie)
}
