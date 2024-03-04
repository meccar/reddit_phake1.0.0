// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: account.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const updateAccount = `-- name: UpdateAccount :one
UPDATE Account
SET
  password = COALESCE($1, password),
  -- password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  -- username = COALESCE(sqlc.narg(username), username),
  is_email_verified = COALESCE($2, is_email_verified)
WHERE
  username = $3
RETURNING id, role, username, password, is_email_verified, created_at
`

type UpdateAccountParams struct {
	Password        pgtype.Text `json:"password"`
	IsEmailVerified pgtype.Bool `json:"is_email_verified"`
	Username        string      `json:"username"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount, arg.Password, arg.IsEmailVerified, arg.Username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
		&i.IsEmailVerified,
		&i.CreatedAt,
	)
	return i, err
}

const authPassword = `-- name: authPassword :one
SELECT password
FROM Account
WHERE username = $1
LIMIT 1
`

func (q *Queries) authPassword(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRow(ctx, authPassword, username)
	var password string
	err := row.Scan(&password)
	return password, err
}

const authUsername = `-- name: authUsername :one
SELECT username
FROM Account
WHERE username = $1
AND is_email_verified = TRUE
LIMIT 1
`

func (q *Queries) authUsername(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRow(ctx, authUsername, username)
	err := row.Scan(&username)
	return username, err
}

const createAccount = `-- name: createAccount :one
INSERT INTO Account (
  id,
  username,
  password,
  role,
  created_at
) VALUES (
  $1,$2,$3,$4,CURRENT_TIMESTAMP
) RETURNING id, role, username, password, is_email_verified, created_at
`

type createAccountParams struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     Userrole  `json:"role"`
}

func (q *Queries) createAccount(ctx context.Context, arg createAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Role,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
		&i.IsEmailVerified,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountIDbyID = `-- name: getAccountIDbyID :one
SELECT id
FROM Account
WHERE id = $1
LIMIT 1
`

func (q *Queries) getAccountIDbyID(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getAccountIDbyID, id)
	err := row.Scan(&id)
	return id, err
}

const getAccountIDbyUsername = `-- name: getAccountIDbyUsername :one
SELECT id
FROM Account
WHERE username = $1
LIMIT 1
`

func (q *Queries) getAccountIDbyUsername(ctx context.Context, username string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getAccountIDbyUsername, username)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getAccountRolebyUsername = `-- name: getAccountRolebyUsername :one
SELECT role
FROM Account
WHERE username = $1
LIMIT 1
`

func (q *Queries) getAccountRolebyUsername(ctx context.Context, username string) (Userrole, error) {
	row := q.db.QueryRow(ctx, getAccountRolebyUsername, username)
	var role Userrole
	err := row.Scan(&role)
	return role, err
}

const loginAccount = `-- name: loginAccount :one
INSERT INTO Account (
  id,
  username,
  password,
  created_at
) VALUES (
  $1,$2,$3,CURRENT_TIMESTAMP
) RETURNING id, role, username, password, is_email_verified, created_at
`

type loginAccountParams struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func (q *Queries) loginAccount(ctx context.Context, arg loginAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, loginAccount, arg.ID, arg.Username, arg.Password)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Username,
		&i.Password,
		&i.IsEmailVerified,
		&i.CreatedAt,
	)
	return i, err
}
