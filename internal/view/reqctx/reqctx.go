package reqctx

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

// AllCtx represents the full context of a request.
//
// ⚠️ SHOULD have with all the other *Ctx structs.
type AllCtx struct {
	Auth AuthCtx

	IsAuthed bool
	UserID   uuid.UUID
}

// AuthCtx represents the authentication values of a user
// in the context of a request.
type AuthCtx struct {
	IsAuthed bool
	User     dbgen.User
}
