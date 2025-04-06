run:
	go run ./cmd/main.go

fresh:
	rm ./internal/data/sql/db.db || true
	sqlc generate
	just run

db:
	sqlite3 ./internal/data/sql/db.db -header
