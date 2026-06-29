package auth

import (
	"shop/internal/api"
	"shop/internal/auth"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestIsLoginUserValid_ValidInput(t *testing.T) {
	loginUser := api.LoginUser{
		Email:    "test@example.com",
		Password: "password123",
	}

	if !auth.IsLoginUserValid(loginUser) {
		t.Error("expected valid login user, got invalid")
	}
}

func TestIsLoginUserValid_EmptyEmail(t *testing.T) {
	loginUser := api.LoginUser{
		Email:    "",
		Password: "password123",
	}

	if auth.IsLoginUserValid(loginUser) {
		t.Error("expected invalid when email is empty")
	}
}

func TestIsLoginUserValid_EmptyPassword(t *testing.T) {
	loginUser := api.LoginUser{
		Email:    "test@example.com",
		Password: "",
	}

	if auth.IsLoginUserValid(loginUser) {
		t.Error("expected invalid when password is empty")
	}
}

func TestIsLoginUserValid_BothEmpty(t *testing.T) {
	loginUser := api.LoginUser{}

	if auth.IsLoginUserValid(loginUser) {
		t.Error("expected invalid when both fields are empty")
	}
}

func TestVerifyPassword_CorrectPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate hash: %v", err)
	}

	if !auth.VerifyPassword(password, hash) {
		t.Error("expected password to match hash")
	}
}

func TestVerifyPassword_WrongPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate hash: %v", err)
	}

	if auth.VerifyPassword("wrongpassword", hash) {
		t.Error("expected wrong password to not match hash")
	}
}

func TestVerifyPassword_EmptyPassword(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("realpassword"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate hash: %v", err)
	}

	if auth.VerifyPassword("", hash) {
		t.Error("expected empty password to not match hash")
	}
}
