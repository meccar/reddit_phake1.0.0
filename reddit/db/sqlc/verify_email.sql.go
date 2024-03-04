// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: verify_email.sql

package db

import (
	"context"
)

const createVerifyEmail = `-- name: CreateVerifyEmail :one
INSERT INTO Verify_email (
    username,
    secret_code
) VALUES (
    $1, $2
) RETURNING username, secret_code, is_used, expires_at, created_at
`

type CreateVerifyEmailParams struct {
	Username   string `json:"username"`
	SecretCode string `json:"secret_code"`
}

func (q *Queries) CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error) {
	row := q.db.QueryRow(ctx, createVerifyEmail, arg.Username, arg.SecretCode)
	var i VerifyEmail
	err := row.Scan(
		&i.Username,
		&i.SecretCode,
		&i.IsUsed,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateVerifyEmail = `-- name: UpdateVerifyEmail :one
UPDATE Verify_email
SET
    is_used = TRUE
WHERE
    username = $1
    AND secret_code = $2
    AND is_used = FALSE
    AND expires_at > now()
RETURNING username, secret_code, is_used, expires_at, created_at
`

type UpdateVerifyEmailParams struct {
	Username   string `json:"username"`
	SecretCode string `json:"secret_code"`
}

func (q *Queries) UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error) {
	row := q.db.QueryRow(ctx, updateVerifyEmail, arg.Username, arg.SecretCode)
	var i VerifyEmail
	err := row.Scan(
		&i.Username,
		&i.SecretCode,
		&i.IsUsed,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
