package restorations

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/integration"
	"github.com/eduardolat/pgbackweb/internal/service/databases"
	"github.com/eduardolat/pgbackweb/internal/service/destinations"
	"github.com/eduardolat/pgbackweb/internal/service/executions"
)

type Service struct {
	dbgen               *dbgen.Queries
	ints                *integration.Integration
	executionsService   *executions.Service
	databasesService    *databases.Service
	destinationsService *destinations.Service
}

func New(
	dbgen *dbgen.Queries, ints *integration.Integration,
	executionsService *executions.Service, databasesService *databases.Service,
	destinationsService *destinations.Service,
) *Service {
	return &Service{
		dbgen:               dbgen,
		ints:                ints,
		executionsService:   executionsService,
		databasesService:    databasesService,
		destinationsService: destinationsService,
	}
}
