// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: login.sql

package db

import (
	"context"
)

const getAccount = `-- name: GetAccount :one
SELECT username, password FROM tb_login
WHERE username = $1 and password=$2
`

type GetAccountParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) GetAccount(ctx context.Context, arg GetAccountParams) (TbLogin, error) {
	row := q.db.QueryRowContext(ctx, getAccount, arg.Username, arg.Password)
	var i TbLogin
	err := row.Scan(&i.Username, &i.Password)
	return i, err
}
