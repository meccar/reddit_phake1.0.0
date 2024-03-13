-- name: createPost :one
INSERT INTO Post (
  id,
  title,
  article,
  picture,
  user_id,
  community_id,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,$6,CURRENT_TIMESTAMP
) RETURNING *;

-- name: getAllPost :many
SELECT *
FROM Post;