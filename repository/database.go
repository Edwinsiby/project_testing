package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=edwin password=acid dbname=sample_project sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// "postgres", "host=sample_project port=5432 user=edwin password=acid dbname=sample_project sslmode=disable"
// "postgres", "host=localhost port=5432 user=edwin password=acid dbname=sample_project sslmode=disable"

// func ConnectTestDB(t *testing.T) (*sql.DB, error) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	return db, nil
// }
