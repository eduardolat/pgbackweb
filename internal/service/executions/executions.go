package executions

import "github.com/eduardolat/pgbackweb/internal/database/dbgen"

type Service struct {
	dbgen *dbgen.Queries
}

func New(dbgen *dbgen.Queries) *Service {
	return &Service{
		dbgen: dbgen,
	}
}
