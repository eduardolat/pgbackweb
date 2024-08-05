package restorations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/google/uuid"
)

type PaginateRestorationsParams struct {
	Page            int
	Limit           int
	ExecutionFilter uuid.NullUUID
	DatabaseFilter  uuid.NullUUID
}

func (s *Service) PaginateRestorations(
	ctx context.Context, params PaginateRestorationsParams,
) (paginateutil.PaginateResponse, []dbgen.RestorationsServicePaginateRestorationsRow, error) {
	page := max(params.Page, 1)
	limit := min(max(params.Limit, 1), 100)

	count, err := s.dbgen.RestorationsServicePaginateRestorationsCount(
		ctx, dbgen.RestorationsServicePaginateRestorationsCountParams{
			ExecutionID: params.ExecutionFilter,
			DatabaseID:  params.DatabaseFilter,
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

	restorations, err := s.dbgen.RestorationsServicePaginateRestorations(
		ctx, dbgen.RestorationsServicePaginateRestorationsParams{
			ExecutionID: params.ExecutionFilter,
			DatabaseID:  params.DatabaseFilter,
			Limit:       int32(params.Limit),
			Offset:      int32(offset),
		},
	)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	return paginateResponse, restorations, nil
}
