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

func (s *Service) ClearSessionCookie(c echo.Context) {
	cookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	c.SetCookie(&cookie)
}

func (s *Service) GetUserFromSessionCookie(c echo.Context) (
	bool, dbgen.AuthServiceGetUserByTokenRow, error,
) {
	ctx := c.Request().Context()

	cookie, err := c.Cookie(sessionCookieName)
	if err != nil {
		return false, dbgen.AuthServiceGetUserByTokenRow{}, err
	}

	if cookie.Value == "" {
		return false, dbgen.AuthServiceGetUserByTokenRow{}, nil
	}

	return s.GetUserByToken(ctx, cookie.Value)
}
