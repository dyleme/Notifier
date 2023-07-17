// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users.sql

package queries

import (
	"context"
)

const addUser = `-- name: AddUser :one
INSERT INTO users (email,
                   nickname,
                   password_hash
                   )
VALUES (
        $1,
        $2,
        $3
       )
RETURNING id
`

type AddUserParams struct {
	Email        string
	Nickname     string
	PasswordHash string
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (int32, error) {
	row := q.db.QueryRow(ctx, addUser, arg.Email, arg.Nickname, arg.PasswordHash)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getLoginParameters = `-- name: GetLoginParameters :one
SELECT id,
       password_hash
FROM users
WHERE email = $1
   OR nickname = $1
`

type GetLoginParametersRow struct {
	ID           int32
	PasswordHash string
}

func (q *Queries) GetLoginParameters(ctx context.Context, authName string) (GetLoginParametersRow, error) {
	row := q.db.QueryRow(ctx, getLoginParameters, authName)
	var i GetLoginParametersRow
	err := row.Scan(&i.ID, &i.PasswordHash)
	return i, err
}
