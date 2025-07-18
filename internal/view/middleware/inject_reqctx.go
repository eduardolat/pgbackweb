package middleware

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/labstack/echo/v4"
	htmx "github.com/nodxdev/nodxgo-htmx"
)

func (m *Middleware) InjectReqctx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqCtx := reqctx.Ctx{
			IsHTMXBoosted: htmx.ServerGetIsBoosted(c.Request().Header),
		}

		found, user, err := m.servs.AuthService.GetUserFromSessionCookie(c)
		if err != nil {
			logger.Error("failed to get user from session cookie", logger.KV{
				"ip":    c.RealIP(),
				"ua":    c.Request().UserAgent(),
				"error": err,
			})
			return c.String(http.StatusInternalServerError, "Internal server error")
		}

		if found {
			reqCtx.IsAuthed = true
			reqCtx.SessionID = user.SessionID
			reqCtx.User = dbgen.User{
				ID:           user.ID,
				Name:         user.Name,
				Email:        user.Email,
				CreatedAt:    user.CreatedAt,
				UpdatedAt:    user.UpdatedAt,
				OidcProvider: user.OidcProvider,
				OidcSubject:  user.OidcSubject,
				Password:     user.Password,
			}
		}

		reqctx.SetCtx(c, reqCtx)
		return next(c)
	}
}
