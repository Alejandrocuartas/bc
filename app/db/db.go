package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("PG_URI"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
