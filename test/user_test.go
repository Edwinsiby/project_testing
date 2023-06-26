package test

import (
	"test/domain"
	"test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	testUser := domain.User{
		UserName: "edwin",
		Password: "pass@123",
	}

	err := repository.GetByName(testUser)
	if err != nil {
		if err.Error() != "username already exists" {
			t.Errorf("Expected error 'username already exists', got '%s'", err.Error())
		}
	} else {
		t.Errorf("Expected error 'username already exists', got no error")
	}

}

func TestCreateUser(t *testing.T) {
	testUser := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	err := repository.CreateUser(testUser)
	assert.NoError(t, err, "Expected no error")

	err = repository.DeleteUser(testUser)
	assert.NoError(t, err, "Expected no error")

	err = repository.GetByName(testUser)
	assert.NoError(t, err, "Expected no error")

}

func TestLogin(t *testing.T) {
	testUser := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	err := repository.Login(testUser)
	assert.Error(t, err, "Expected error")
	assert.EqualError(t, err, "invalid username or password", "Expected error message 'invalid username or password'")
}
