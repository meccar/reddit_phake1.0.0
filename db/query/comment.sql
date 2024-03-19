-- name: createComment :one
INSERT INTO Comment (
  id,
  post_id,
  user_id,
  text,
  upvotes,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,CURRENT_TIMESTAMP
) RETURNING *;

-- name: GetCommentFromPost :many
SELECT * FROM Comment
WHERE post_id = $1
ORDER BY created_at DESC;


