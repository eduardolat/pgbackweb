package integration

import (
	"github.com/eduardolat/pgbackweb/internal/integration/postgres"
	"github.com/eduardolat/pgbackweb/internal/integration/s3"
)

type Integration struct {
	PGClient *postgres.Client
	S3Client *s3.Client
}

func New() *Integration {
	pgClient := postgres.New()
	s3Client := s3.New()

	return &Integration{
		PGClient: pgClient,
		S3Client: s3Client,
	}
}
