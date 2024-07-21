package backups

import "github.com/google/uuid"

func (s *Service) jobRemove(backupID uuid.UUID) error {
	return s.cr.RemoveJob(backupID)
}
