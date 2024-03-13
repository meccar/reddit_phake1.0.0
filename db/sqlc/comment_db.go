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
