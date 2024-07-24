package auth

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) GetUserSessions(
	ctx context.Context, userID uuid.UUID,
) ([]dbgen.Session, error) {
	return s.dbgen.AuthServiceGetUserSessions(ctx, userID)
}
