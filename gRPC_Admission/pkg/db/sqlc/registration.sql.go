// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: registration.sql

package db

import (
	"context"
	"database/sql"
)

const getRegistration = `-- name: GetRegistration :one
SELECT id, username, hashed_password, full_name, email, updated_at, created_at FROM registration
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetRegistration(ctx context.Context, username string) (Registration, error) {
	row := q.db.QueryRowContext(ctx, getRegistration, username)
	var i Registration
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const registration = `-- name: Registration :one
INSERT INTO registration (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING id, username, hashed_password, full_name, email, updated_at, created_at
`

type RegistrationParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) Registration(ctx context.Context, arg RegistrationParams) (Registration, error) {
	row := q.db.QueryRowContext(ctx, registration,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i Registration
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateRegistration = `-- name: UpdateRegistration :one
UPDATE registration
SET
  hashed_password = COALESCE($1, hashed_password),
  updated_at = COALESCE($2, updated_at),
  full_name = COALESCE($3, full_name),
  email = COALESCE($4, email)
WHERE
  username = $5
RETURNING id, username, hashed_password, full_name, email, updated_at, created_at
`

type UpdateRegistrationParams struct {
	HashedPassword sql.NullString `json:"hashed_password"`
	UpdatedAt      sql.NullTime   `json:"updated_at"`
	FullName       sql.NullString `json:"full_name"`
	Email          sql.NullString `json:"email"`
	Username       string         `json:"username"`
}

func (q *Queries) UpdateRegistration(ctx context.Context, arg UpdateRegistrationParams) (Registration, error) {
	row := q.db.QueryRowContext(ctx, updateRegistration,
		arg.HashedPassword,
		arg.UpdatedAt,
		arg.FullName,
		arg.Email,
		arg.Username,
	)
	var i Registration
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}