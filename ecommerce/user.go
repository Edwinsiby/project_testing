package ecommerce

import (
	"database/sql"
	"fmt"
	"test/repository"
)

var db *sql.DB
var err error

func init() {
	db, err = repository.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
}

type User struct {
	Username string
	Password string
}

type UserManager struct {
	users []User
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (um *UserManager) Signup(username, password string) error {
	query := "SELECT COUNT(*) FROM users WHERE username = $1"
	var count int
	err = db.QueryRow(query, username).Scan(&count)
	if err != nil {
		return nil
	}
	if count > 0 {
		return fmt.Errorf("username already exists")
	}

	query = "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = db.Exec(query, username, password)
	if err != nil {
		return err
	}

	return nil
}

func (um *UserManager) Login(username, password string) error {
	query := "SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2"
	var count int
	err := db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}
