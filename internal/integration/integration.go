package integration

import (
	"github.com/eduardolat/pgbackweb/internal/integration/pgdump"
	"github.com/eduardolat/pgbackweb/internal/integration/s3"
)

type Integration struct {
	PGDumpClient *pgdump.Client
	S3Client     *s3.Client
}

func New() *Integration {
	pgdumpClient := pgdump.New()
	s3Client := s3.New()

	return &Integration{
		PGDumpClient: pgdumpClient,
		S3Client:     s3Client,
	}
}
