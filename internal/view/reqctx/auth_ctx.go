package reqctx

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/labstack/echo/v4"
)

// SetAuthCtx inserts the authentication values of
// a user into a echo request context.
func SetAuthCtx(c echo.Context, auth AuthCtx) {
	c.Set("isAuthed", auth.IsAuthed)
	c.Set("user", auth.User)
}

// GetAuthCtx returns the authentication values of
// a user from a echo request context.
func GetAuthCtx(c echo.Context) AuthCtx {
	var isAuthed bool
	var user dbgen.User

	if ia, ok := c.Get("isAuthed").(bool); ok {
		isAuthed = ia
	}
	if au, ok := c.Get("user").(dbgen.User); ok {
		user = au
	}

	return AuthCtx{
		IsAuthed: isAuthed,
		User:     user,
	}
}
