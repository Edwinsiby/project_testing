package ecommerce

import "fmt"

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
	// Check if the user already exists
	for _, user := range um.users {
		if user.Username == username {
			return fmt.Errorf("username already exists")
		}
	}

	// Create a new user
	newUser := User{
		Username: username,
		Password: password,
	}
	um.users = append(um.users, newUser)

	return nil
}

func (um *UserManager) Login(username, password string) error {
	for _, user := range um.users {
		if user.Username == username && user.Password == password {
			return nil
		}
	}

	return fmt.Errorf("invalid username or password")
}
