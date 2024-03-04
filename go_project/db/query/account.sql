-- name: createAccount :one
INSERT INTO Account (
  id,
  username,
  password,
  role,
  created_at
) VALUES (
  $1,$2,bcrypt($3),$4,CURRENT_TIMESTAMP
) RETURNING *;

-- name: loginAccount :one
INSERT INTO Account (
  id,
  username,
  password,
  created_at
) VALUES (
  $1,$2,bcrypt($3),CURRENT_TIMESTAMP
) RETURNING *;

-- name: authUsername :one
SELECT username
FROM Account
WHERE username = $1
LIMIT 1;

-- name: authPassword :one
SELECT password
FROM Account
WHERE username = $1
LIMIT 1;