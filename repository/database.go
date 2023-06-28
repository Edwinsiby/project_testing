package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=rohith password=892328 dbname=test sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectTestDB() (*sql.DB, error) {
	testDB, err := sql.Open("postgres", "host=localhost port=5432 user=edwin password=acid dbname=test_databse sslmode=disable")
	if err != nil {
		return nil, err
	}
	return testDB, nil
}
