// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package data

import (
	"context"

	"github.com/google/uuid"
)

const addUser = `-- name: AddUser :exec
INSERT INTO users (id, name, email, password_hash)
VALUES (?, ?, ?, ?)
`

type AddUserParams struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) error {
	_, err := q.db.ExecContext(ctx, addUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.PasswordHash,
	)
	return err
}

const getByEmail = `-- name: GetByEmail :one
SELECT id, name, email, password_hash FROM users
WHERE email = ? LIMIT 1
`

func (q *Queries) GetByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
	)
	return i, err
}
