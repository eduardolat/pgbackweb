package webhooks

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
)

type PaginateWebhooksParams struct {
	Page  int
	Limit int
}

func (s *Service) PaginateWebhooks(
	ctx context.Context, params PaginateWebhooksParams,
) (paginateutil.PaginateResponse, []dbgen.Webhook, error) {
	page := max(params.Page, 1)
	limit := min(max(params.Limit, 1), 100)

	count, err := s.dbgen.WebhooksServicePaginateWebhooksCount(ctx)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	paginateParams := paginateutil.PaginateParams{
		Page:  page,
		Limit: limit,
	}
	offset := paginateutil.CreateOffsetFromParams(paginateParams)
	paginateResponse := paginateutil.CreatePaginateResponse(paginateParams, int(count))

	webhooks, err := s.dbgen.WebhooksServicePaginateWebhooks(
		ctx, dbgen.WebhooksServicePaginateWebhooksParams{
			Limit:  int32(params.Limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	return paginateResponse, webhooks, nil
}
