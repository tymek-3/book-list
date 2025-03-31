package data

import (
	"context"
	"database/sql"
	"log"
	"os"
)

func NewDB(logger *log.Logger) (*Queries, error) {
	ctx := context.Background()

	dbPath := "./internal/data/sql/db.db"
	_, err := os.Stat(dbPath)
	dbExists := !os.IsNotExist(err)

	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if !dbExists {
		schema, err := os.ReadFile("./internal/data/sql/schema.sql")
		if err != nil {
			return nil, err
		}

		if _, err = conn.ExecContext(ctx, string(schema)); err != nil {
			return nil, err
		}
		logger.Println("database created")
	}

	db := New(conn)
	logger.Println("database initialized")
	return db, nil
}
