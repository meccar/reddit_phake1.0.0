package db

import (
	"context"

	"github.com/google/uuid"
)

func (h *Handlers) GetAccountbyID(ctx context.Context, id uuid.UUID) ([]GetAccountbyIDRow, error) {
	return h.Queries.GetAccountbyID(ctx, id)
}

func (h *Handlers) GetAccountIDbyUsername(ctx context.Context, username string) (uuid.UUID, error) {
	return h.Queries.GetAccountIDbyUsername(ctx, username)
}
