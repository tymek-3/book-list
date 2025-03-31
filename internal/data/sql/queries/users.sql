-- name: AddUser :exec
INSERT INTO users (id, name, email, password_hash)
VALUES (?, ?, ?, ?);

-- name: GetByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;
