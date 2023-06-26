package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"test/domain"
)

var db *sql.DB
var err error

func init() {
	db, err = ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
}

func GetByName(user domain.User) error {
	query := "SELECT COUNT(*) FROM users WHERE username = $1"
	var count int
	err = db.QueryRow(query, user.UserName).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}
	return nil
}

func CreateUser(user domain.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = db.Exec(query, user.UserName, user.Password)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func Login(user domain.User) error {
	query := "SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2"
	var count int
	err = db.QueryRow(query, user.UserName, user.Password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}
