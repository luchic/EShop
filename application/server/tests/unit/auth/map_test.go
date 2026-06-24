package auth

import (
	"bytes"
	"shop/internal/api"
	"shop/internal/auth"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestMapRegisterUserToUser_SetsCorrectFields(t *testing.T) {
	password := "password123"

	registerUser := api.RegisterUser{
		FirstName:  "John",
		SecondName: "Doe",
		Email:      "john@example.com",
		Password:   password,
	}

	user, err := auth.MapRegisterUserToUser(registerUser)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.FirstName != registerUser.FirstName {
		t.Errorf("FirstName = %q, want %q", user.FirstName, registerUser.FirstName)
	}
	if user.SecondName != registerUser.SecondName {
		t.Errorf("SecondName = %q, want %q", user.SecondName, registerUser.SecondName)
	}
	if user.Email != registerUser.Email {
		t.Errorf("Email = %q, want %q", user.Email, registerUser.Email)
	}
	if user.Role != "user" {
		t.Errorf("Role = %q, want %q", user.Role, "user")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	if !bytes.Equal(user.Password, hashPassword) {
		t.Errorf("Hash passwords not equal")
	}
}

func TestMapRegisterUserToUser_HashesPassword(t *testing.T) {
	password := "password123"
	registerUser := api.RegisterUser{
		FirstName:  "John",
		SecondName: "Doe",
		Email:      "john@example.com",
		Password:   password,
	}

	user, err := auth.MapRegisterUserToUser(registerUser)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(user.Password) == password {
		t.Error("password should be hashed, not stored as plain text")
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		t.Error("hashed password should match original password")
	}
}

func TestMapRegisterUserToUser_DifferentInputsProduceDifferentHashes(t *testing.T) {
	registerUser := api.RegisterUser{
		FirstName: "John",
		Email:     "john@example.com",
		Password:  "password123",
	}

	user1, err := auth.MapRegisterUserToUser(registerUser)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	user2, err := auth.MapRegisterUserToUser(registerUser)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(user1.Password) == string(user2.Password) {
		t.Error("same password should produce different hashes (bcrypt uses random salt)")
	}
}
