-- name: createCommunity :one
INSERT INTO Community (
  id,
  community_name,
  photo,
  created_at
) VALUES (
  $1,$2,$3,CURRENT_TIMESTAMP
) RETURNING *;

-- name: SearchCommunityName :many
SELECT id, community_name
FROM Community
WHERE community_name like $1;

-- name: GetCommunityIDbyName :one
SELECT id FROM Community
WHERE community_name = $1;

-- name: GetCommunitybyID :many
SELECT * FROM Community
WHERE id = $1;

-- -- name: GetAllCommunity :many
-- SELECT * FROM Community
-- ORDER BY created_at DESC;