package db

import (
	"context"

	"github.com/google/uuid"
)

type CreateCommentTxParams struct {
	createCommentParams
}

type CreateCommentTxResult struct {
	Comment *Comment
}

func (h *Handlers) CreateCommentTx(ctx context.Context, arg CreateCommentTxParams) (CreateCommentTxResult, error) {
	var result CreateCommentTxResult

	err := h.execTx(ctx, func(q *Queries) error {
		var err error

		ranID, _ := uuid.NewRandom()

		// // Submit the form to the database
		params := createCommentParams{
			ID:      ranID,
			PostID:  arg.PostID,
			UserID:  arg.UserID,
			Text:    arg.Text,
			Upvotes: arg.Upvotes,
		}

		Comment, err := q.createComment(ctx, params)

		if err != nil {
			return err
		}

		result.Comment = &Comment
		return err
	})
	return result, err
}

func (h *Handlers) GetCommentFromPost(ctx context.Context, postID uuid.UUID) ([]Comment, error) {
	return h.Queries.GetCommentFromPost(ctx, postID)
}

func (h *Handlers) FilterPostHavingMostComments(ctx context.Context) ([]FilterPostHavingMostCommentsRow, error) {
	return h.Queries.FilterPostHavingMostComments(ctx)
}

func (h *Handlers) GetCommentFromUser(ctx context.Context, userID uuid.UUID) ([]Comment, error) {
	return h.Queries.GetCommentFromUser(ctx, userID)
}

func (h *Handlers) SortCommentASC(ctx context.Context) ([]Comment, error) {
	return h.Queries.SortCommentASC(ctx)
}

func (h *Handlers) SortCommentDESC(ctx context.Context) ([]Comment, error) {
	return h.Queries.SortCommentDESC(ctx)
}

func (h *Handlers) SortCommentByUpvotes(ctx context.Context) ([]Comment, error) {
	return h.Queries.SortCommentByUpvotes(ctx)
}
