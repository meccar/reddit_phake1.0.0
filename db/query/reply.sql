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

-- name: GetReplyFromComment :many
SELECT * FROM Reply
WHERE comment_id = $1
ORDER BY created_at DESC;