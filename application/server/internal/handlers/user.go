package handlers

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"shop/internal/api"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const SESSION_DURATION = 30 * time.Minute

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

	if !isLoginUserValid(loginUser) {
		http.Error(w, "Email and password requiered", http.StatusBadRequest)
	}

	user, err := h.repository.GetUserByEmail(loginUser.Email)
	if err != nil {
		http.Error(w, "Invalid password or Login", http.StatusUnauthorized)
		return
	}

	if !verifyPassword(loginUser.Password, user.Password) {
		http.Error(w, "Invalid password or Login", http.StatusUnauthorized)
		return
	}

	sessionId, expiresAt, err := h.createSession(&user)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
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

func (h *Handler) createSession(user *api.User) (string, time.Time, error) {
	sessionId, err := generateSessionId()
	if err != nil {
		return "", time.Time{}, err
	}

	timeNow := time.Now()
	expiredAt := timeNow.Add(SESSION_DURATION)

	sessionData := api.SessionData{
		UserID:    user.Id,
		Email:     user.Email,
		CreatedAt: timeNow,
		ExpiresAt: expiredAt,
	}

	sessionJSON, err := json.Marshal(sessionData)
	if err != nil {
		return "", time.Time{}, err
	}

	err = h.redis.Set(
		context.Background(),
		sessionId,
		sessionJSON,
		SESSION_DURATION).Err()
	if err != nil {
		return "", time.Time{}, err
	}

	return sessionId, expiredAt, nil
}

func generateSessionId() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func isLoginUserValid(loginUser api.LoginUser) bool {
	_ = loginUser
	return true
}

func verifyPassword(login_password string, actual_hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(actual_hash, []byte(login_password))
	return err == nil
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
