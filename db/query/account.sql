-- name: createAccount :one
INSERT INTO Account (
  id,
  username,
  password,
  role,
  photo,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,CURRENT_TIMESTAMP
) RETURNING *;

-- name: loginAccount :one
INSERT INTO Account (
  id,
  username,
  password,
  created_at
) VALUES (
  $1,$2,$3,CURRENT_TIMESTAMP
) RETURNING *;

-- name: authUsername :one
SELECT username
FROM Account
WHERE username = $1
AND is_email_verified = TRUE
LIMIT 1;

-- name: authPassword :one
SELECT password
FROM Account
WHERE username = $1
LIMIT 1;

-- name: getAccountIDbyID :one
SELECT id
FROM Account
WHERE id = $1
LIMIT 1;

-- name: GetAccountIDbyUsername :one
SELECT id
FROM Account
WHERE username = $1
LIMIT 1;

-- name: getAccountRolebyUsername :one
SELECT role
FROM Account
WHERE username = $1
LIMIT 1;

-- name: GetAccountbyID :many
SELECT   id, username, photo
FROM Account
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE Account
SET
  password = COALESCE(sqlc.narg(password), password),
  -- password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  -- username = COALESCE(sqlc.narg(username), username),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
  username = sqlc.arg(username)
RETURNING *;