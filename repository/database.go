package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func ConnectDB() (DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=edwin password=acid dbname=sample_project sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
