-- name: createCommunity :one
INSERT INTO Community (
  id,
  community_name,
  photo,
  created_at
) VALUES (
  $1,$2,$3,CURRENT_TIMESTAMP
) RETURNING *;