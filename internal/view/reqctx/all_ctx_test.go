package reqctx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAllCtxFuncs(t *testing.T) {
	testUser := dbgen.User{
		ID:    uuid.New(),
		Email: "user@example.com",
		Name:  "John",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("Create authentication values in context", func(t *testing.T) {
		SetAuthCtx(c, AuthCtx{
			IsAuthed: true,
			User:     testUser,
		})

		ctx := GetAllCtx(c)

		assert.True(t, ctx.Auth.IsAuthed)
		assert.Equal(t, testUser, ctx.Auth.User)
		assert.Equal(t, testUser.Email, ctx.Auth.User.Email)
		assert.Equal(t, testUser.ID, ctx.UserID)
		assert.Equal(t, testUser.ID, ctx.Auth.User.ID)
	})
}
