-- name: Registration :one
INSERT INTO registration (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetRegistration :one
SELECT * FROM registration
WHERE username = $1 LIMIT 1;

-- name: UpdateRegistration :one
UPDATE registration
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  email = COALESCE(sqlc.narg(email), email)
WHERE
  username = sqlc.arg(username)
RETURNING *;
