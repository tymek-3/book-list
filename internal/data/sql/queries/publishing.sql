-- name: AuthorsAdd :exec
INSERT INTO authors (id, full_name)
VALUES (?, ?);

-- name: PublishersAdd :exec
INSERT INTO publishers (id, name)
VALUES (?, ?);

-- name: TypesAdd :exec
INSERT INTO types (id, name)
VALUES (?, ?);
