// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const register = `-- name: Register :one
INSERT INTO users(
    id,
    username, 
    password,
    fullname,
    gender,
    avt,
    role_id
) VALUES (
    $1, $2, $3, $4, $5, $6, 2
) RETURNING id, username, password, fullname, gender, avt, role_id, created_at
`

type RegisterParams struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Fullname string    `json:"fullname"`
	Gender   int32     `json:"gender"`
	Avt      string    `json:"avt"`
}

func (q *Queries) Register(ctx context.Context, arg RegisterParams) (User, error) {
	row := q.db.QueryRowContext(ctx, register,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Fullname,
		arg.Gender,
		arg.Avt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Fullname,
		&i.Gender,
		&i.Avt,
		&i.RoleID,
		&i.CreatedAt,
	)
	return i, err
}
