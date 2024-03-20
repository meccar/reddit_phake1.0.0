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

-- name: GetCommentFromUser :many
SELECT * FROM Comment
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: SortCommentByUpvotes :many
SELECT * FROM Comment
ORDER BY upvotes DESC;

-- name: SortCommentDESC :many
SELECT * FROM Comment
ORDER BY created_at DESC;

-- name: SortCommentASC :many
SELECT * FROM Comment
ORDER BY created_at ASC;

-- name: FilterPostHavingMostComments :many
SELECT post_id, COUNT(*) AS comment_count
FROM Comment
GROUP BY post_id
ORDER BY comment_count DESC
LIMIT 1;




