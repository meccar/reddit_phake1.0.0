package db

import (
	"context"

	"github.com/google/uuid"
)

type CreatePostTxParams struct {
	createPostParams
	CommunityName string
}

type CreatePostTxResult struct {
	Post *Post
}

func (h *Handlers) CreatePostTx(ctx context.Context, arg CreatePostTxParams) (CreatePostTxResult, error) {
	var result CreatePostTxResult

	err := h.execTx(ctx, func(q *Queries) error {

		ranID, err := uuid.NewRandom()

		// // Submit the form to the database
		params := createPostParams{
			ID:          ranID,
			Title:       arg.Title,
			Article:     arg.Article,
			Picture:     arg.Picture,
			UserID:      arg.UserID,
			CommunityID: arg.CommunityID,
		}

		Post, err := q.createPost(ctx, params)

		if err != nil {
			return err
		}

		result.Post = &Post
		return err
	})
	return result, err
}

func (h *Handlers) GetAllPosts(ctx context.Context) ([]Post, error) {
	return h.Queries.GetAllPost(ctx)
}
