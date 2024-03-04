// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: session.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: createSession :many
INSERT INTO Session(
  token,
  role,
  issuer,
  subject,
  exp,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,CURRENT_TIMESTAMP
) RETURNING id, token, role, issuer, subject, exp, created_at
`

type createSessionParams struct {
	Token   string           `json:"token"`
	Role    string           `json:"role"`
	Issuer  string           `json:"issuer"`
	Subject string           `json:"subject"`
	Exp     pgtype.Timestamp `json:"exp"`
}

func (q *Queries) createSession(ctx context.Context, arg createSessionParams) ([]Session, error) {
	rows, err := q.db.Query(ctx, createSession,
		arg.Token,
		arg.Role,
		arg.Issuer,
		arg.Subject,
		arg.Exp,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Token,
			&i.Role,
			&i.Issuer,
			&i.Subject,
			&i.Exp,
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

const removeSession = `-- name: removeSession :exec
DELETE FROM Session
WHERE token = $1
`

func (q *Queries) removeSession(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, removeSession, token)
	return err
}
