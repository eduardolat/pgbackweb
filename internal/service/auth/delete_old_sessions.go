package auth

import (
	"context"
	"time"
)

const maxSessionAge = time.Hour * 12

func (s *Service) DeleteOldSessions(ctx context.Context) error {
	dateThreshold := time.Now().Add(-maxSessionAge)
	return s.dbgen.AuthServiceDeleteOldSessions(ctx, dateThreshold)
}
