package restorations

import "context"

func (s *Service) GetRestorationsQty(ctx context.Context) (int64, error) {
	return s.dbgen.RestorationsServiceGetRestorationsQty(ctx)
}
