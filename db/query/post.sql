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

-- name: GetAllPost :many
SELECT * FROM Post
ORDER BY created_at DESC;

-- name: GetPostByCommunity :many
SELECT * FROM Post
WHERE community_id = $1
ORDER BY created_at DESC;

-- name: GetPostByUser :many
SELECT * FROM Post
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetPostbyID :one
SELECT * FROM Post
WHERE id = $1;

-- name: SortPostByUpvotesDESC :many
SELECT * FROM Post
ORDER BY upvotes DESC;

-- name: SortPostByUpvotesASC :many
SELECT * FROM Post
ORDER BY upvotes ASC;

-- name: SortPostByDay :many
SELECT * FROM Post
WHERE DATE_PART('day', created_at) = $1;

-- name: SortPostByMonth :many
SELECT * FROM Post
WHERE DATE_PART('month', created_at) = $1;

-- name: SortPostByYear :many
SELECT * FROM Post
WHERE DATE_PART('year', created_at) = $1;

    
