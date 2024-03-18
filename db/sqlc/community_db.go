package db

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

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

		imagePath := "https://tafviet.com/wp-content/uploads/2024/03/community-picture.jpg"

		// Make an HTTP GET request to fetch the image data
		resp, err := http.Get(imagePath)
		if err != nil {
			fmt.Println("Error fetching image:", err)
			return err
		}
		defer resp.Body.Close()

		// Read the response body
		imageBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading image data:", err)
			return err
		}

		// Encode the image to base64
		base64Encoded := base64.StdEncoding.EncodeToString(imageBytes)

		// If the photo is empty, set it to the base64-encoded image data
		if len(arg.Photo) == 0 {
			arg.Photo = []byte(base64Encoded)
		}

		// // Submit the form to the database
		params := createCommunityParams{
			ID:            ranID,
			CommunityName: arg.CommunityName,
			Photo:         arg.Photo,
		}

		Community, err := q.createCommunity(ctx, params)
		fmt.Println("CreateCommunityTx: ", Community)

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

				result, err := h.CreateCommunityTx(ctx, arg)

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
