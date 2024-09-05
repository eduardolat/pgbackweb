package reqctx

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	ctxKey = "PGBackWebCTX"
)

// Ctx represents the values passed through a single request context.
type Ctx struct {
	IsHTMXBoosted bool
	IsAuthed      bool
	SessionID     uuid.UUID
	User          dbgen.User
}

// SetCtx inserts values into the Echo request context.
func SetCtx(c echo.Context, ctx Ctx) {
	c.Set(ctxKey, ctx)
}

// GetCtx retrieves values from the Echo request context.
func GetCtx(c echo.Context) Ctx {
	ctx, ok := c.Get(ctxKey).(Ctx)
	if !ok {
		return Ctx{}
	}
	return ctx
}
