package db

import (
	"context"
	"time"

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

func (h *Handlers) GetPostbyID(ctx context.Context, id uuid.UUID) (Post, error) {
	return h.Queries.GetPostbyID(ctx, id)
}

func (h *Handlers) SortPostByYear(ctx context.Context, createdAt time.Time) ([]Post, error) {
	return h.Queries.SortPostByYear(ctx, createdAt)
}

func (h *Handlers) SortPostByMonth(ctx context.Context, createdAt time.Time) ([]Post, error) {
	return h.Queries.SortPostByMonth(ctx, createdAt)
}

func (h *Handlers) SortPostByDay(ctx context.Context, createdAt time.Time) ([]Post, error) {
	return h.Queries.SortPostByDay(ctx, createdAt)
}

func (h *Handlers) SortPostByUpvotesASC(ctx context.Context) ([]Post, error) {
	return h.Queries.SortPostByUpvotesASC(ctx)
}

func (h *Handlers) SortPostByUpvotesDESC(ctx context.Context) ([]Post, error) {
	return h.Queries.SortPostByUpvotesDESC(ctx)
}
