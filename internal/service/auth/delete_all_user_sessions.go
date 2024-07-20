package auth

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteAllUserSessions(
	ctx context.Context, userID uuid.UUID,
) error {
	return s.dbgen.AuthServiceDeleteAllUserSessions(ctx, userID)
}
