package main_test

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"test/ecommerce"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcommerceIntegration(t *testing.T) {
	// Build the application
	appPath := buildApplication(t)

	// Run the application and capture the output
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

	// Validate the output
	expectedOutput := "Cart Items:\n- Product 1\n- Product 2\n\nAvailable Products:\n- Product 1 ($10.99)\n- Product 2 ($19.99)\n"
	assert.Equal(t, expectedOutput, stdout.String())
	// Test user signup and login
	userManager := ecommerce.NewUserManager()
	err = userManager.Signup("john", "password")
	assert.NoError(t, err, "failed to sign up user")
	err = userManager.Login("john", "password")
	assert.NoError(t, err, "failed to log in with correct credentials")
	err = userManager.Login("john", "wrongpassword")
	assert.Error(t, err, "expected error when logging in with incorrect password")
	err = userManager.Login("jane", "password")
	assert.Error(t, err, "expected error when logging in with incorrect username")
}

func buildApplication(t *testing.T) string {
	appPath, err := filepath.Abs("test")
	assert.NoError(t, err, "failed to get absolute path")

	// Build the application
	cmd := exec.Command("go", "build", "-o", appPath, ".")
	err = cmd.Run()
	assert.NoError(t, err, "failed to build the application")

	return appPath
}
