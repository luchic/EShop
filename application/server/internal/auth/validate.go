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

func VerifyPassword(loginPassowrd string, actualHash []byte) bool {
	err := bcrypt.CompareHashAndPassword(actualHash, []byte(loginPassowrd))
	return err == nil
}
