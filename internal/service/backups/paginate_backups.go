package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
)

type PaginateBackupsParams struct {
	Page  int
	Limit int
}

func (s *Service) PaginateBackups(
	ctx context.Context, params PaginateBackupsParams,
) (paginateutil.PaginateResponse, []dbgen.BackupsServicePaginateBackupsRow, error) {
	page := max(params.Page, 1)
	limit := min(max(params.Limit, 1), 100)

	count, err := s.dbgen.BackupsServicePaginateBackupsCount(ctx)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	paginateParams := paginateutil.PaginateParams{
		Page:  page,
		Limit: limit,
	}
	offset := paginateutil.CreateOffsetFromParams(paginateParams)
	paginateResponse := paginateutil.CreatePaginateResponse(paginateParams, int(count))

	backups, err := s.dbgen.BackupsServicePaginateBackups(
		ctx, dbgen.BackupsServicePaginateBackupsParams{
			Limit:  int32(params.Limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	return paginateResponse, backups, nil
}
