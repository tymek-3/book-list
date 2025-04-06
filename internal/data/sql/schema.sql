-- TODO: createdat updatedat

CREATE TABLE books (
	id uuid PRIMARY KEY,
	name TEXT NOT null,
	score REAL,
	publication_date TEXT,
	type_id uuid NOT NULL,
	author_id uuid,
	publisher_id uuid,

	FOREIGN KEY(type_id) REFERENCES types(id),
	FOREIGN KEY(author_id) REFERENCES authors(id),
	FOREIGN KEY(publisher_id) REFERENCES publishers(id)
);

CREATE TABLE types(
	id uuid PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE publishers(
	id uuid PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE authors(
	id uuid PRIMARY KEY,
	full_name TEXT NOT NULL
);

CREATE TABLE users(
	email TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	role TEXT NOT NULL,
	role_set_by TEXT,
	password_hash TEXT NOT NULL
);

CREATE TABLE list_item(
	user_email TEXT NOT NULL,
	book_id uuid NOT NULL,
	status INTEGER NOT NULL,

	PRIMARY KEY(user_email, book_id)
);
