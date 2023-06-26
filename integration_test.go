package main_test

import (
	"bytes"
	"database/sql"
	"os/exec"
	"path/filepath"
	"test/ecommerce"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB
var err error
var userManager *ecommerce.UserManager

// func init() {
// 	testDB, err = mocks.ConnectTestDb()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func TestEcommerceIntegration(t *testing.T) {
	appPath := buildApplication(t)

	cmd := exec.Command(appPath)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		t.Errorf("Failed to run the application: %v", err)
		t.Logf("Stdout: %s", stdout.String())
		t.Logf("Stderr: %s", stderr.String())
	}

	expectedOutput := "Cart Items:\n- Product 1\n- Product 2\n\nAvailable Products:\n- Product 1 ($10.99)\n- Product 2 ($19.99)\nusername already exists\nuser logedin\n"
	assert.Equal(t, expectedOutput, stdout.String())

	err = userManager.Login("john", "password")
	assert.NoError(t, err, "user logged in successfully")

	err = userManager.Login("john", "password")
	assert.NoError(t, err, "user logedin")

	err = userManager.Login("john", "wrongpassword")
	assert.Error(t, err, "expected error when logging in with incorrect password")

	err = userManager.Login("jane", "password")
	assert.Error(t, err, "expected error when logging in with incorrect username")
}

func buildApplication(t *testing.T) string {
	appPath, err := filepath.Abs("test")
	assert.NoError(t, err, "failed to get absolute path")

	cmd := exec.Command("go", "build", "-o", appPath, ".")
	err = cmd.Run()
	assert.NoError(t, err, "failed to build the application")

	return appPath
}
