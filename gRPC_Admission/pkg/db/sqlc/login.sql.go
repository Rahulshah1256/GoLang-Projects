// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: login.sql

package db

import (
	"context"
)

const login = `-- name: Login :one
SELECT username, hashed_password FROM login
WHERE username = $1 and hashed_password=$2
`

type LoginParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) Login(ctx context.Context, arg LoginParams) (Login, error) {
	row := q.db.QueryRowContext(ctx, login, arg.Username, arg.HashedPassword)
	var i Login
	err := row.Scan(&i.Username, &i.HashedPassword)
	return i, err
}
