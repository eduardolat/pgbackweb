package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) GetAllBackups(
	ctx context.Context,
) ([]dbgen.Backup, error) {
	return s.dbgen.BackupsServiceGetAllBackups(ctx)
}
