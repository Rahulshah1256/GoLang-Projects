-- name: Admission :one
INSERT INTO admission (
  full_name,
  gender,
  dob,
  address,
  phone_no,
  previous_company,
  designation,
  skills
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetAdmission :one
SELECT * FROM admission
WHERE full_name = $1 LIMIT 1;

-- name: UpdateAdmission :one
UPDATE admission
SET
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  gender = COALESCE(sqlc.narg(gender), gender),
  dob = COALESCE(sqlc.narg(dob), dob),
  address = COALESCE(sqlc.narg(address), address),
  phone_no = COALESCE(sqlc.narg(phone_no), phone_no),
  previous_company = COALESCE(sqlc.narg(previous_company), previous_company),
  designation = COALESCE(sqlc.narg(designation), designation),
  skills = COALESCE(sqlc.narg(skills), skills)
WHERE
  full_name = sqlc.arg(full_name)
RETURNING *;
