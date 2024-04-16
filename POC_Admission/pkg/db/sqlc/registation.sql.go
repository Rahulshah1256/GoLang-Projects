// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: registation.sql

package db

import (
	"context"
	"database/sql"
)

const getData = `-- name: GetData :one
SELECT id, username, full_name, email, password, create_at, updated_at FROM tb_registration
WHERE username = $1 and password=$2
`

type GetDataParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) GetData(ctx context.Context, arg GetDataParams) (TbRegistration, error) {
	row := q.db.QueryRowContext(ctx, getData, arg.Username, arg.Password)
	var i TbRegistration
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreateAt,
		&i.UpdatedAt,
	)
	return i, err
}

const registration = `-- name: Registration :one
INSERT INTO tb_registration (
  username, full_name,email,password,create_at,updated_at
) VALUES (
  $1, $2,$3,$4,$5,$6
)
RETURNING id, username, full_name, email, password, create_at, updated_at
`

type RegistrationParams struct {
	Username  string         `json:"username"`
	FullName  string         `json:"full_name"`
	Email     sql.NullString `json:"email"`
	Password  string         `json:"password"`
	CreateAt  sql.NullTime   `json:"create_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

func (q *Queries) Registration(ctx context.Context, arg RegistrationParams) (TbRegistration, error) {
	row := q.db.QueryRowContext(ctx, registration,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.Password,
		arg.CreateAt,
		arg.UpdatedAt,
	)
	var i TbRegistration
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreateAt,
		&i.UpdatedAt,
	)
	return i, err
}