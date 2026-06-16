package auth

import (
	"shop/internal/api"

	"golang.org/x/crypto/bcrypt"
)

func IsLoginUserValid(loginUser api.LoginUser) bool {
	if loginUser.Email == "" || loginUser.Password == "" {
		return false
	}
	return true
}

func VerifyPassword(login_password string, actual_hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(actual_hash, []byte(login_password))
	return err == nil
}
