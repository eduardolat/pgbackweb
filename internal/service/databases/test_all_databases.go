package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"golang.org/x/sync/errgroup"
)

func (s *Service) TestAllDatabases() {
	ctx := context.Background()

	databases, err := s.GetAllDatabases(ctx)
	if err != nil {
		logger.Error(
			"error getting all databases to test them", logger.KV{"error": err},
		)
		return
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(5)

	for _, db := range databases {
		db := db
		eg.Go(func() error {
			err := s.TestDatabaseAndStoreResult(ctx, db.ID)
			if err != nil {
				logger.Error(
					"error testing database",
					logger.KV{"database_id": db.ID, "error": err},
				)
			}
			return nil
		})
	}

	_ = eg.Wait()
	logger.Info("all databases tested")

}
