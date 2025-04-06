-- name: BookAdd :exec
INSERT INTO books (id, name, score, publication_date, author_id, publisher_id)
VALUES (?, ?, ?, ?, ?, ?);

-- name: BookGetFullById :one
SELECT
b.id,
b.name,
b.score,
b.publication_date,

a.id AS author_id,
a.full_name AS author_full_name,

t.id AS type_id,
t.name AS type_name,

p.id AS publisher_id,
p.name AS publisher_name

FROM books AS b
LEFT JOIN authors AS a ON b.author_id = a.id
LEFT JOIN publishers AS p ON b.publisher_id = p.id
LEFT JOIN types AS t ON b.type_id = t.id
WHERE b.id = ? LIMIT 1;

-- name: BookFullSearch :many
SELECT
b.id,
b.name,
b.score,
b.publication_date,

a.id AS author_id,
a.full_name AS author_full_name,

t.id AS type_id,
t.name AS type_name,

p.id AS publisher_id,
p.name AS publisher_name

FROM books AS b
LEFT JOIN authors AS a ON b.author_id = a.id
LEFT JOIN publishers AS p ON b.publisher_id = p.id
LEFT JOIN types AS t ON b.type_id = t.id
WHERE b.name LIKE '%' || ? || '%'
ORDER BY b.score
LIMIT 20;

-- name: BookFullAll :many
SELECT
b.id,
b.name,
b.score,
b.publication_date,

a.id AS author_id,
a.full_name AS author_full_name,

t.id AS type_id,
t.name AS type_name,

p.id AS publisher_id,
p.name AS publisher_name

FROM books AS b
LEFT JOIN authors AS a ON b.author_id = a.id
LEFT JOIN publishers AS p ON b.publisher_id = p.id
LEFT JOIN types AS t ON b.type_id = t.id
ORDER BY b.score;
