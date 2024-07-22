package auth

import (
	"context"
	"time"

	"github.com/eduardolat/pgbackweb/internal/logger"
)

const maxSessionAge = time.Hour * 12

func (s *Service) DeleteOldSessions() {
	ctx := context.Background()
	dateThreshold := time.Now().Add(-maxSessionAge)

	err := s.dbgen.AuthServiceDeleteOldSessions(ctx, dateThreshold)
	if err != nil {
		logger.Error(
			"error deleting old sessions", logger.KV{"error": err},
		)
		return
	}

	logger.Info("old sessions deleted")
}
