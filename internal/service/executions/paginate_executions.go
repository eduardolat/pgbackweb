package executions

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/google/uuid"
)

type PaginateExecutionsParams struct {
	Page              int
	Limit             int
	DatabaseFilter    uuid.NullUUID
	DestinationFilter uuid.NullUUID
	BackupFilter      uuid.NullUUID
}

func (s *Service) PaginateExecutions(
	ctx context.Context, params PaginateExecutionsParams,
) (paginateutil.PaginateResponse, []dbgen.Execution, error) {
	page := max(params.Page, 1)
	limit := min(max(params.Limit, 1), 100)

	count, err := s.dbgen.ExecutionsServicePaginateExecutionsCount(
		ctx, dbgen.ExecutionsServicePaginateExecutionsCountParams{
			BackupID:      params.BackupFilter,
			DatabaseID:    params.DatabaseFilter,
			DestinationID: params.DestinationFilter,
		},
	)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	paginateParams := paginateutil.PaginateParams{
		Page:  page,
		Limit: limit,
	}
	offset := paginateutil.CreateOffsetFromParams(paginateParams)
	paginateResponse := paginateutil.CreatePaginateResponse(paginateParams, int(count))

	executions, err := s.dbgen.ExecutionsServicePaginateExecutions(
		ctx, dbgen.ExecutionsServicePaginateExecutionsParams{
			BackupID:      params.BackupFilter,
			DatabaseID:    params.DatabaseFilter,
			DestinationID: params.DestinationFilter,
			Limit:         int32(params.Limit),
			Offset:        int32(offset),
		},
	)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	return paginateResponse, executions, nil
}
