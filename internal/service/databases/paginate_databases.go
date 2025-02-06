package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
)

type PaginateDatabasesParams struct {
	Page  int
	Limit int
}

func (s *Service) PaginateDatabases(
	ctx context.Context, params PaginateDatabasesParams,
) (paginateutil.PaginateResponse, []dbgen.DatabasesServicePaginateDatabasesRow, error) {
	page := max(params.Page, 1)
	limit := min(max(params.Limit, 1), 100)

	count, err := s.dbgen.DatabasesServicePaginateDatabasesCount(ctx)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	paginateParams := paginateutil.PaginateParams{
		Page:  page,
		Limit: limit,
	}
	offset := paginateutil.CreateOffsetFromParams(paginateParams)
	paginateResponse := paginateutil.CreatePaginateResponse(paginateParams, int(count))

	databases, err := s.dbgen.DatabasesServicePaginateDatabases(
		ctx, dbgen.DatabasesServicePaginateDatabasesParams{
			EncryptionKey: s.env.PBW_ENCRYPTION_KEY,
			Limit:         int32(params.Limit),
			Offset:        int32(offset),
		},
	)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	return paginateResponse, databases, nil
}
