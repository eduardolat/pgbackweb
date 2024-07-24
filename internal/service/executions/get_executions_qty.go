package executions

import "context"

func (s *Service) GetExecutionsQty(ctx context.Context) (int64, error) {
	return s.dbgen.ExecutionsServiceGetExecutionsQty(ctx)
}
