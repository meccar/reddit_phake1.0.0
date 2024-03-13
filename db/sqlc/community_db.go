package db

import (
	"context"

	"github.com/google/uuid"
)

type CreateCommunityTxParams struct {
	createCommunityParams
}

type CreateCommunityTxResult struct {
	Community *Community
}

func (h *Handlers) CreateCommunityTx(ctx context.Context, arg CreateCommunityTxParams) (CreateCommunityTxResult, error) {
	var result CreateCommunityTxResult

	err := h.execTx(ctx, func(q *Queries) error {

		ranID, err := uuid.NewRandom()

		// // Submit the form to the database
		params := createCommunityParams{
			ID:            ranID,
			CommunityName: arg.CommunityName,
			Photo:         arg.Photo,
		}

		Community, err := q.createCommunity(ctx, params)

		if err != nil {
			return err
		}

		result.Community = &Community
		return err
	})
	return result, err
}

func (h *Handlers) SearchCommunity(ctx context.Context, communityName string) ([]SearchCommunityNameRow, error) {
	return h.Queries.SearchCommunityName(ctx, communityName)
}
