package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
)

type PaginateDestinationsParams struct {
	Page  int
	Limit int
}

func (s *Service) PaginateDestinations(
	ctx context.Context, params PaginateDestinationsParams,
) (paginateutil.PaginateResponse, []dbgen.DestinationsServicePaginateDestinationsRow, error) {
	page := max(params.Page, 1)
	limit := max(params.Limit, 100)

	count, err := s.dbgen.DestinationsServicePaginateDestinationsCount(ctx)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	paginateParams := paginateutil.PaginateParams{
		Page:  page,
		Limit: limit,
	}
	offset := paginateutil.CreateOffsetFromParams(paginateParams)
	paginateResponse := paginateutil.CreatePaginateResponse(paginateParams, int(count))

	destinations, err := s.dbgen.DestinationsServicePaginateDestinations(
		ctx, dbgen.DestinationsServicePaginateDestinationsParams{
			EncryptionKey: *s.env.PBW_ENCRYPTION_KEY,
			Limit:         int32(params.Limit),
			Offset:        int32(offset),
		},
	)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	return paginateResponse, destinations, nil
}
