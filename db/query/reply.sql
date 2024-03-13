-- name: createReply :one
INSERT INTO Reply (
  id,
  comment_id,
  user_id,
  text,
  upvotes,
  created_at
) VALUES (
  $1,$2,$3,$4,$5,CURRENT_TIMESTAMP
) RETURNING *;