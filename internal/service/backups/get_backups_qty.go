package backups

import "context"

func (s *Service) GetBackupsQty(ctx context.Context) (int64, error) {
	return s.dbgen.BackupsServiceGetBackupsQty(ctx)
}
