-- name: UserAdd :exec
INSERT INTO users (id, name, email, role, password_hash)
VALUES (?, ?, ?, ?, ?);

-- name: UserGetByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;
