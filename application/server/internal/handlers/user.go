package handlers

import (
	"encoding/json"
	"net/http"
	"shop/internal/api"
	"shop/internal/auth"
)

// handleRegisterUser godoc
// @Summary      Register a new user
// @Tags         users
// @Accept       json
// @Produce      plain
// @Param        body  body      api.RegisterUser  true  "Registration payload"
// @Success      201
// @Failure      400  {string}  string  "Bad request"
// @Router       /user/register [post]
func (h *Handler) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registerUser api.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&registerUser); err != nil {
		http.Error(w, "BBad Request", http.StatusBadRequest)
		return
	}

	user, err := auth.MapRegisterUserToUser(registerUser)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	if err := h.repository.CreateUser(user); err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// handleLoginUser godoc
// @Summary      Login
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      api.LoginUser      true  "Login payload"
// @Success      200   {object}  api.LoginResponse
// @Failure      400   {string}  string  "Bad request"
// @Failure      401   {string}  string  "Unauthorized"
// @Failure      500   {string}  string  "Internal server error"
// @Router       /user/login [post]
func (h *Handler) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginUser api.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Couldn't decode json body", http.StatusBadRequest)
		return
	}

	if !auth.IsLoginUserValid(loginUser) {
		http.Error(w, "Email and password requiered", http.StatusBadRequest)
		return
	}

	user, err := h.repository.GetUserByEmail(loginUser.Email)
	if err != nil {
		http.Error(w, "Invalid password or Login", http.StatusUnauthorized)
		return
	}

	if !auth.VerifyPassword(loginUser.Password, user.Password) {
		http.Error(w, "Invalid password or Login", http.StatusUnauthorized)
		return
	}

	sessionId, expiresAt, err := h.auth.CreateSession(&user)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	response := api.LoginResponse{
		UserId:    user.Id,
		SessionId: sessionId,
		Email:     user.Email,
		ExpiresAt: expiresAt.Unix(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) handleLogOut(w http.ResponseWriter, r *http.Request) {

}
