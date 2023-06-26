package mocks

import "database/sql"

var testDB *sql.DB

func connectTestDb() (*sql.DB, error) {
	testDB, err := sql.Open("postgres", "host=localhost port=5432 user=edwin password=acid dbname=test_database sslmode=disable")
	if err != nil {
		return nil, err
	}
	return testDB, nil
}
