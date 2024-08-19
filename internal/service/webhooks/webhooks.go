package webhooks

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/orsinium-labs/enum"
)

type webhook = enum.Member[string]

var (
	eventTypeDatabaseHealthy   = webhook{Value: "database_healthy"}
	eventTypeDatabaseUnhealthy = webhook{Value: "database_unhealthy"}

	eventTypeDestinationHealthy   = webhook{Value: "destination_healthy"}
	eventTypeDestinationUnhealthy = webhook{Value: "destination_unhealthy"}

	eventTypeExecutionSuccess = webhook{Value: "execution_success"}
	eventTypeExecutionFailed  = webhook{Value: "execution_failed"}
)

type Service struct {
	dbgen *dbgen.Queries
}

func New(
	dbgen *dbgen.Queries,
) *Service {
	return &Service{
		dbgen: dbgen,
	}
}
