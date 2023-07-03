package repository

import (
	"database/sql"
	"log"
	"test/utils"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load  env", err)
	}
	db, err := sql.Open(config.DB, config.DNS)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DB="postgres",
// DNS="host=172.17.0.2 port=5432 user=edwin password=acid dbname=sample_project sslmode=disable"

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
