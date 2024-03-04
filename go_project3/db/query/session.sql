-- name: createSession :one
INSERT INTO Session(
  id,
  username,
  role,
  expires_at,
  created_at
) VALUES (
  $1,$2,$3,$4,CURRENT_TIMESTAMP
) RETURNING *;

-- name: deleteSession :exec
DELETE FROM Session
WHERE username = $1;

-- name: getSessionIDbyID :one
SELECT id
FROM Session
WHERE id = $1
LIMIT 1;

-- name: getAllSessionID :many
SELECT id
FROM Session;