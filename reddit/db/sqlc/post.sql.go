// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: post.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPost = `-- name: createPost :one
INSERT INTO Post (
  id,
  title,
  article,
  picture,
  created_at
) VALUES (
  $1,$2,$3,$4,CURRENT_TIMESTAMP
) RETURNING id, title, article, picture, created_at
`

type createPostParams struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Article string    `json:"article"`
	Picture []byte    `json:"picture"`
}

func (q *Queries) createPost(ctx context.Context, arg createPostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.ID,
		arg.Title,
		arg.Article,
		arg.Picture,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Article,
		&i.Picture,
		&i.CreatedAt,
	)
	return i, err
}

const getAllPost = `-- name: getAllPost :many
SELECT id, title, article, picture, created_at
FROM Post
`

func (q *Queries) getAllPost(ctx context.Context) ([]Post, error) {
	rows, err := q.db.Query(ctx, getAllPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Article,
			&i.Picture,
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
