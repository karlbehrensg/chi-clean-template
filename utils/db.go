package utils

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

func NewPostgresClient(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	slog.Info("Connected to Postgres database")
	return db, nil
}
