package databases

import "context"

func (s *Service) GetDatabasesQty(ctx context.Context) (int64, error) {
	return s.dbgen.DatabasesServiceGetDatabasesQty(ctx)
}
