-- name: Registration :one
INSERT INTO tb_registration (
  username, full_name,email,password,create_at,updated_at
) VALUES (
  $1, $2,$3,$4,$5,$6
)
RETURNING *;

-- name: GetData :one
SELECT * FROM tb_registration
WHERE username = $1 and password=$2;