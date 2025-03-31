run:
	go run ./cmd/main.go

fresh:
	rm ./internal/data/sql/db.db
	sqlc generate
	just run
