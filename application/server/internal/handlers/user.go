package handlers

import (
	"encoding/json"
	"net/http"
	"shop/internal/api"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registerUser api.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&registerUser); err != nil {
		http.Error(w, "Couldn't decode json body", http.StatusBadRequest)
		return
	}

	user := mapRegisterUserToUser(registerUser)

	if err := h.repository.CreateUsesr(user); err != nil {
		http.Error(w, "Internal Error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Implementation isn't so good.
func mapRegisterUserToUser(registerUser api.RegisterUser) api.User {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)

	user := api.User{
		FirstName:  registerUser.FirstName,
		SecondName: registerUser.SecondName,
		Email:      registerUser.Email,
		Password:   hashPassword,
	}
	return user
}
