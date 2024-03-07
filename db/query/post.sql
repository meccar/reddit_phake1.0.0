-- name: createPost :one
INSERT INTO Post (
  id,
  title,
  article,
  picture,
  created_at
) VALUES (
  $1,$2,$3,$4,CURRENT_TIMESTAMP
) RETURNING *;

-- name: getAllPost :many
SELECT *
FROM Post;