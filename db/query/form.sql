-- name: submitForm :one
INSERT INTO Form(
  id,
  viewer_name,
  email,
  phone,
  created_at
) VALUES (
  $1,$2,$3,$4,CURRENT_TIMESTAMP
) RETURNING *;

-- name: getFormsID :one
SELECT id
FROM Form
WHERE id = $1
LIMIT 1;