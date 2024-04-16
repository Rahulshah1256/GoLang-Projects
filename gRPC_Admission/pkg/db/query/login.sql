-- name: Login :one
SELECT * FROM login
WHERE username = $1 and hashed_password=$2;