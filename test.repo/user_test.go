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

//Mock

// func TestGetByNameMock(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("failed to create mock: %v", err)
// 	}
// 	defer db.Close()

// 	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
// 	mock.ExpectQuery("SELECT COUNT(*) FROM users WHERE username = ?").
// 		WithArgs("testuser").
// 		WillReturnRows(rows)

// 	user := domain.User{UserName: "testuser"}
// 	err = repository.GetByName(user)

// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("unfulfilled expectations: %v", err)
// 	}
// }
