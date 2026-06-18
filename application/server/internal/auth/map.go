package auth

import (
	"shop/internal/api"

	"golang.org/x/crypto/bcrypt"
)

// Implementation isn't so good.
func MapRegisterUserToUser(registerUser api.RegisterUser) (api.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return api.User{}, err
	}

	user := api.User{
		Role:       "user",
		FirstName:  registerUser.FirstName,
		SecondName: registerUser.SecondName,
		Email:      registerUser.Email,
		Password:   hashPassword,
	}
	return user, nil
}
