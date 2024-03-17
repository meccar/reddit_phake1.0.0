package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

		if len(arg.Photo.String) == 0 || arg.Photo.String == "" {
			arg.Photo = pgtype.Text{String: "https://tafviet.com/wp-content/uploads/2024/03/community-picture.jpg"}
		}

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

func (h *Handlers) GetCommunitybyID(ctx context.Context, id uuid.UUID) ([]Community, error) {
	return h.Queries.GetCommunitybyID(ctx, id)
}

func (h *Handlers) GetCommunityIDbyName(ctx context.Context, communityName string) (uuid.UUID, error) {
	fmt.Println("\n GetCommunityIDbyName communityName:", communityName)

	id, err := h.Queries.GetCommunityIDbyName(ctx, communityName)
	fmt.Println("\n GetCommunityIDbyName id:", id)
	fmt.Println("\n GetCommunityIDbyName err:", err)

	if err != nil {
		if err.Error() == "no rows in result set" {
			// Create a channel to receive the community ID and error
			ch := make(chan struct {
				ID  uuid.UUID
				Err error
			})

			// Asynchronously create the community
			go func() {
				arg := CreateCommunityTxParams{
					createCommunityParams: createCommunityParams{
						CommunityName: communityName,
					},
				}
				fmt.Println("\n GetCommunityIDbyName arg:", arg)

				result, err := h.CreateCommunityTx(ctx, arg)
				fmt.Println("\n GetCommunityIDbyName result:", result)

				// Send the community ID and error through the channel
				ch <- struct {
					ID  uuid.UUID
					Err error
				}{result.Community.ID, err}

			}()

			// Receive the community ID and error from the channel
			result := <-ch
			if result.Err != nil {
				return uuid.UUID{}, result.Err
			}
			return result.ID, nil

		}
		return uuid.UUID{}, err
	}
	return id, nil
}
