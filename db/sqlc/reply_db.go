package db

import (
	"context"

	"github.com/google/uuid"
)

type CreateReplyTxParams struct {
	createReplyParams
}

type CreateReplyTxResult struct {
	Reply *Reply
}

func (h *Handlers) createReplyTx(ctx context.Context, arg CreateReplyTxParams) (CreateReplyTxResult, error) {
	var result CreateReplyTxResult

	err := h.execTx(ctx, func(q *Queries) error {
		var err error

		ranID, _ := uuid.NewRandom()

		// // Submit the form to the database
		params := createReplyParams{
			ID:        ranID,
			CommentID: arg.CommentID,
			UserID:    arg.UserID,
			Text:      arg.Text,
			Upvotes:   arg.Upvotes,
		}

		Reply, err := q.createReply(ctx, params)

		if err != nil {
			return err
		}

		result.Reply = &Reply
		return err
	})
	return result, err
}
