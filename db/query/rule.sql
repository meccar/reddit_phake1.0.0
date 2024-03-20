-- name: createRule :one
INSERT INTO Rule (
  id,
  community_id,
  title,
  description
) VALUES (
  $1,$2,$3,$4
) RETURNING *;

-- name: GetRuleFromCommunity :many
SELECT * FROM Rule
WHERE community_id = $1;