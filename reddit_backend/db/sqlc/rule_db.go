package db

import (
	"context"

	"github.com/google/uuid"
)

func (h *Handlers) GetRuleFromCommunity(ctx context.Context, communityID uuid.UUID) ([]Rule, error) {
	return h.Queries.GetRuleFromCommunity(ctx, communityID)
}
