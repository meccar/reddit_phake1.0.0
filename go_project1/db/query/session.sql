-- name: createSession :many
INSERT INTO Session(
  token,
  role,
  issuer,
  subject,
  exp,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,CURRENT_TIMESTAMP
) RETURNING *;

-- name: removeSession :exec
DELETE FROM Session
WHERE token = $1;