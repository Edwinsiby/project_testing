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

// func ConnectTestDB(t *testing.T) (*sql.DB, error) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	return db, nil
// }
