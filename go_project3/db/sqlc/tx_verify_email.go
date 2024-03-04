package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

)


type CreateVerifyEmailTxResult struct {
	VerifyEmail *VerifyEmail
}

func (h *Handlers) VerifyEmail(ctx context.Context, arg *CreateVerifyEmailParams) (CreateVerifyEmailTxResult, error) {
	var result CreateVerifyEmailTxResult

	err := h.execTx(ctx, func(q *Queries) error {
		var err error

		verifier, err := q.CreateVerifyEmail(ctx, *arg)
		if err != nil {
			return err
		}

		_, err = q.UpdateAccount(ctx, UpdateAccountParams{
			Username: verifier.Username,
			IsEmailVerified: pgtype.Bool{
				Bool:  true,
				Valid: true,
			},
		})
		if err != nil {
			return err
		}

		_, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			Username: verifier.Username,
			SecretCode: verifier.SecretCode,
		})
		if err != nil {
			return err
		}

		result.VerifyEmail = &verifier
		return err
	})

	return result, err
}