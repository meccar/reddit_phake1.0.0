// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: community.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createCommunity = `-- name: createCommunity :one
INSERT INTO Community (
  id,
  community_name,
  created_at
) VALUES (
  $1,$2,CURRENT_TIMESTAMP
) RETURNING id, community_name, photo, created_at
`

type createCommunityParams struct {
	ID            uuid.UUID `json:"id"`
	CommunityName string    `json:"community_name"`
}

func (q *Queries) createCommunity(ctx context.Context, arg createCommunityParams) (Community, error) {
	row := q.db.QueryRow(ctx, createCommunity, arg.ID, arg.CommunityName)
	var i Community
	err := row.Scan(
		&i.ID,
		&i.CommunityName,
		&i.Photo,
		&i.CreatedAt,
	)
	return i, err
}