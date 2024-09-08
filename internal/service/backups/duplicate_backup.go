package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) DuplicateBackup(
	ctx context.Context, backupID uuid.UUID,
) (dbgen.Backup, error) {
	return s.dbgen.BackupsServiceDuplicateBackup(ctx, backupID)
}
