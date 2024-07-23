package reqctx

import (
	"github.com/labstack/echo/v4"
)

// GetAllCtx returns ALL the values from an
// echo request context.
//
// It includes AuthCtx, TenantCtx, etc.
func GetAllCtx(c echo.Context) AllCtx {
	authCtx := GetAuthCtx(c)

	return AllCtx{
		Auth: authCtx,

		IsAuthed: authCtx.IsAuthed,
		UserID:   authCtx.User.ID,
	}
}
