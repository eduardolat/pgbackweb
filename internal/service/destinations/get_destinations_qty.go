package destinations

import "context"

func (s *Service) GetDestinationsQty(ctx context.Context) (int64, error) {
	return s.dbgen.DestinationsServiceGetDestinationsQty(ctx)
}
