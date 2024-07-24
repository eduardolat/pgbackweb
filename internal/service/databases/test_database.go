package databases

import (
	"context"
	"fmt"
)

func (s *Service) TestDatabase(
	ctx context.Context, version, connString string,
) error {
	pgVersion, err := s.ints.PGDumpClient.ParseVersion(version)
	if err != nil {
		return fmt.Errorf("error parsing PostgreSQL version: %w", err)
	}

	err = s.ints.PGDumpClient.Ping(pgVersion, connString)
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	return nil
}
