// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: reply.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getReplyFromComment = `-- name: GetReplyFromComment :many
SELECT id, comment_id, user_id, text, upvotes, created_at FROM Reply
WHERE comment_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetReplyFromComment(ctx context.Context, commentID uuid.UUID) ([]Reply, error) {
	rows, err := q.db.Query(ctx, getReplyFromComment, commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Reply{}
	for rows.Next() {
		var i Reply
		if err := rows.Scan(
			&i.ID,
			&i.CommentID,
			&i.UserID,
			&i.Text,
			&i.Upvotes,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createReply = `-- name: createReply :one
INSERT INTO Reply (
  id,
  comment_id,
  user_id,
  text,
  upvotes,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,CURRENT_TIMESTAMP
) RETURNING id, comment_id, user_id, text, upvotes, created_at
`

type createReplyParams struct {
	ID        uuid.UUID `json:"id"`
	CommentID uuid.UUID `json:"comment_id"`
	UserID    uuid.UUID `json:"user_id"`
	Text      string    `json:"text"`
	Upvotes   int32     `json:"upvotes"`
}

func (q *Queries) createReply(ctx context.Context, arg createReplyParams) (Reply, error) {
	row := q.db.QueryRow(ctx, createReply,
		arg.ID,
		arg.CommentID,
		arg.UserID,
		arg.Text,
		arg.Upvotes,
	)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.CommentID,
		&i.UserID,
		&i.Text,
		&i.Upvotes,
		&i.CreatedAt,
	)
	return i, err
}
