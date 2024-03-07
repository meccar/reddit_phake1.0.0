package db

import (
	"context"

	"github.com/google/uuid"
)

type CreatePostTxParams struct {
	createPostParams
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
			ID:      ranID,
			Title:   arg.Title,
			Article: arg.Article,
			Picture: arg.Picture,
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

func (h *Handlers) GetPosts(ctx context.Context) ([]Post, error) {
	return h.Queries.getAllPost(ctx)
}
