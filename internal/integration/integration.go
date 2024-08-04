package integration

import (
	"github.com/eduardolat/pgbackweb/internal/integration/postgres"
	"github.com/eduardolat/pgbackweb/internal/integration/storage"
)

type Integration struct {
	PGClient      *postgres.Client
	StorageClient *storage.Client
}

func New() *Integration {
	pgClient := postgres.New()
	storageClient := storage.New()

	return &Integration{
		PGClient:      pgClient,
		StorageClient: storageClient,
	}
}
