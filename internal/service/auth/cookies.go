package auth

import (
	"errors"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/labstack/echo/v4"
)

const (
	sessionCookieName = "pbw_session"
)

func (s *Service) SetSessionCookie(c echo.Context, token string) {
	cookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    token,
		MaxAge:   int(maxSessionAge.Seconds()),
		HttpOnly: true,
		Secure:   true, // Force HTTPS
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	c.SetCookie(&cookie)
}

func (s *Service) ClearSessionCookie(c echo.Context) {
	cookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true, // Force HTTPS
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	c.SetCookie(&cookie)
}

func (s *Service) GetUserFromSessionCookie(c echo.Context) (
	bool, dbgen.AuthServiceGetUserByTokenRow, error,
) {
	ctx := c.Request().Context()

	cookie, err := c.Cookie(sessionCookieName)
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		return false, dbgen.AuthServiceGetUserByTokenRow{}, nil
	}
	if err != nil {
		return false, dbgen.AuthServiceGetUserByTokenRow{}, err
	}

	if cookie.Value == "" {
		return false, dbgen.AuthServiceGetUserByTokenRow{}, nil
	}

	return s.GetUserByToken(ctx, cookie.Value)
}
