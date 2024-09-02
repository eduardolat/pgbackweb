package webhooks

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/paginateutil"
	"github.com/google/uuid"
)

type PaginateWebhookExecutionsParams struct {
	WebhookID uuid.UUID
	Page      int
	Limit     int
}

func (s *Service) PaginateWebhookExecutions(
	ctx context.Context, params PaginateWebhookExecutionsParams,
) (paginateutil.PaginateResponse, []dbgen.WebhookResult, error) {
	page := max(params.Page, 1)
	limit := min(max(params.Limit, 1), 100)

	count, err := s.dbgen.WebhooksServicePaginateWebhookExecutionsCount(
		ctx, params.WebhookID,
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

	webhookResults, err := s.dbgen.WebhooksServicePaginateWebhookExecutions(
		ctx, dbgen.WebhooksServicePaginateWebhookExecutionsParams{
			WebhookID: params.WebhookID,
			Limit:     int32(params.Limit),
			Offset:    int32(offset),
		},
	)
	if err != nil {
		return paginateutil.PaginateResponse{}, nil, err
	}

	return paginateResponse, webhookResults, nil
}
