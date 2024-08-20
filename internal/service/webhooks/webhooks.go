package webhooks

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/orsinium-labs/enum"
)

type (
	eventType     = enum.Member[eventTypeData]
	eventTypeData struct {
		Key  string
		Name string
	}
)

var (
	EventTypeDatabaseHealthy = eventType{
		Value: eventTypeData{Key: "database_healthy", Name: "Database healthy"},
	}
	EventTypeDatabaseUnhealthy = eventType{
		Value: eventTypeData{Key: "database_unhealthy", Name: "Database unhealthy"},
	}

	EventTypeDestinationHealthy = eventType{
		Value: eventTypeData{Key: "destination_healthy", Name: "Destination healthy"},
	}
	EventTypeDestinationUnhealthy = eventType{
		Value: eventTypeData{Key: "destination_unhealthy", Name: "Destination unhealthy"},
	}

	EventTypeExecutionSuccess = eventType{
		Value: eventTypeData{Key: "execution_success", Name: "Execution success"},
	}
	EventTypeExecutionFailed = eventType{
		Value: eventTypeData{Key: "execution_failed", Name: "Execution failed"},
	}
)

var FullEventTypes = map[string]string{
	EventTypeDatabaseHealthy.Value.Key:      EventTypeDatabaseHealthy.Value.Name,
	EventTypeDatabaseUnhealthy.Value.Key:    EventTypeDatabaseUnhealthy.Value.Name,
	EventTypeDestinationHealthy.Value.Key:   EventTypeDestinationHealthy.Value.Name,
	EventTypeDestinationUnhealthy.Value.Key: EventTypeDestinationUnhealthy.Value.Name,
	EventTypeExecutionSuccess.Value.Key:     EventTypeExecutionSuccess.Value.Name,
	EventTypeExecutionFailed.Value.Key:      EventTypeExecutionFailed.Value.Name,
}

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
