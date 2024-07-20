package integration

import "github.com/eduardolat/pgbackweb/internal/integration/pgdump"

type Integration struct {
	PGDumpClient *pgdump.Client
}

func New() *Integration {
	pgdumpClient := pgdump.New()

	return &Integration{
		PGDumpClient: pgdumpClient,
	}
}
