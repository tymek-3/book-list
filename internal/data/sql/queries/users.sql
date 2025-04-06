-- name: UserAdd :exec
INSERT INTO users (name, email, role, password_hash)
VALUES (?, ?, ?, ?);

-- name: UserUpdate :exec
UPDATE users
SET email = ?, name = ?, role = ?, role_set_by = ?, password_hash = ?
WHERE email = ?;

-- name: UserGetByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

