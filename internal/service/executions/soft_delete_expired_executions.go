package executions

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/logger"
)

func (s *Service) SoftDeleteExpiredExecutions() {
	ctx := context.Background()

	expiredExecutions, err := s.dbgen.ExecutionsServiceGetExpiredExecutions(ctx)
	if err != nil {
		logger.Error(
			"error soft deleting expired executions",
			logger.KV{"error": err},
		)
		return
	}

	for _, execution := range expiredExecutions {
		if err := s.SoftDeleteExecution(ctx, execution.ID); err != nil {
			logger.Error(
				"error soft deleting expired executions",
				logger.KV{"id": execution.ID.String(), "error": err},
			)
			return
		}
	}

	logger.Info("expired executions soft deleted")
}
