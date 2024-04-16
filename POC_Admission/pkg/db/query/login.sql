-- name: GetAccount :one
SELECT * FROM tb_login
WHERE username = $1 and password=$2;