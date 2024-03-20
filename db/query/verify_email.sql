-- name: CreateVerifyEmail :one
INSERT INTO Verify_email (
    username,
    secret_code
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE Verify_email
SET
    is_used = TRUE
WHERE
    username = @username
    AND secret_code = @secret_code
    AND is_used = FALSE
    AND expires_at > now()
RETURNING *;