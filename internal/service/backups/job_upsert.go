package backups

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) jobUpsert(
	backupID uuid.UUID, timeZone string, cronExpression string,
) error {
	return s.cr.UpsertJob(
		backupID, timeZone, cronExpression,
		s.executionsService.RunExecution, context.Background(), backupID,
	)
}
