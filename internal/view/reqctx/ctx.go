package reqctx

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/labstack/echo/v4"
)

// Context keys to avoid typos
const (
	isAuthedKey = "isAuthed"
	userKey     = "user"
)

// Ctx represents the values passed through a single request context.
type Ctx struct {
	IsAuthed bool
	User     dbgen.User
}

// SetCtx inserts values into the Echo request context.
func SetCtx(c echo.Context, ctx Ctx) {
	c.Set(isAuthedKey, ctx.IsAuthed)
	c.Set(userKey, ctx.User)
}

// GetCtx retrieves values from the Echo request context.
func GetCtx(c echo.Context) Ctx {
	var isAuthed bool
	var user dbgen.User

	if ia, ok := c.Get(isAuthedKey).(bool); ok {
		isAuthed = ia
	}
	if au, ok := c.Get(userKey).(dbgen.User); ok {
		user = au
	}

	return Ctx{
		IsAuthed: isAuthed,
		User:     user,
	}
}
