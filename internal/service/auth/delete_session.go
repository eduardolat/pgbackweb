package auth

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteSession(
	ctx context.Context, sessionID uuid.UUID,
) error {
	return s.dbgen.AuthServiceDeleteSession(ctx, sessionID)
}
