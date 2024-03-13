-- name: createCommunity :one
INSERT INTO Community (
  id,
  community_name,
  created_at
) VALUES (
  $1,$2,CURRENT_TIMESTAMP
) RETURNING *;