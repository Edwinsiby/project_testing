package ecommerce_test

import (
	"testing"

	"test/ecommerce"

	"github.com/stretchr/testify/assert"
)

func TestUserManager_Signup(t *testing.T) {
	userManager := ecommerce.NewUserManager()

	err := userManager.Signup("john", "password")
	assert.Error(t, err, "failed to sign up user")

	err = userManager.Signup("john", "password")
	assert.Error(t, err, "expected error when signing up with existing username")
}

func TestUserManager_Login(t *testing.T) {
	userManager := ecommerce.NewUserManager()

	err := userManager.Signup("john", "password")
	assert.Error(t, err, "failed to sign up user")

	err = userManager.Login("john", "password")
	assert.NoError(t, err, "failed to log in with correct credentials")

	err = userManager.Login("john", "wrongpassword")
	assert.Error(t, err, "expected error when logging in with incorrect password")

	err = userManager.Login("jane", "password")
	assert.Error(t, err, "expected error when logging in with incorrect username")
}
