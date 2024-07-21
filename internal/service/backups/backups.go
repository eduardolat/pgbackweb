package backups

import (
	"github.com/eduardolat/pgbackweb/internal/cron"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/service/executions"
)

type Service struct {
	dbgen             *dbgen.Queries
	cr                *cron.Cron
	executionsService *executions.Service
}

func New(
	dbgen *dbgen.Queries,
	cr *cron.Cron,
	executionsService *executions.Service,
) *Service {
	return &Service{
		dbgen:             dbgen,
		cr:                cr,
		executionsService: executionsService,
	}
}
