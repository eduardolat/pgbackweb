package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"golang.org/x/sync/errgroup"
)

func (s *Service) TestAllDestinations() {
	ctx := context.Background()

	destinations, err := s.GetAllDestinations(ctx)
	if err != nil {
		logger.Error(
			"error getting all destinations to test them", logger.KV{"error": err},
		)
		return
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(5)

	for _, dest := range destinations {
		dest := dest
		eg.Go(func() error {
			err := s.TestDestinationAndStoreResult(ctx, dest.ID)
			if err != nil {
				logger.Error(
					"error testing destination",
					logger.KV{"destination_id": dest.ID, "error": err},
				)
			}
			return nil
		})
	}

	_ = eg.Wait()
	logger.Info("all destinations tested")
}
