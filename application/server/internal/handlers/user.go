package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"shop/internal/api"
	"shop/internal/auth"
	"shop/internal/services"
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
	// I'm not realy sure, but this messages could be bad. It better use
	// Some logging system to make it.
	// I will just it leave. I want to reaturn nothing for now.
	ctx := r.Context()
	requestId := services.GetRequestId(ctx)
	h.logger.Info("Start method handleRegisterUser.", slog.String("request_id", requestId))
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registerUser api.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&registerUser); err != nil {
		h.logger.Error(
			"Failed to decode json in handleRegisterUser.",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := auth.MapRegisterUserToUser(registerUser)
	if err != nil {
		h.logger.Error(
			"Failed to map models in handleRegisterUser.",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	if err := h.repository.CreateUser(user); err != nil {
		h.logger.Error(
			"Failed to create a new user in handleRegisterUser.",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	sessionId, _, err := h.auth.CreateSession(&user)
	if err != nil {
		h.logger.Error(
			"Failed to create a session in handleRegisterUser.",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:     auth.SESSION_ID_KEY,
		Value:    sessionId,
		Path:     "/",
		HttpOnly: true,
	})
	w.WriteHeader(http.StatusCreated)
	h.logger.Info("Finished method handleRegisterUser.", slog.String("request_id", requestId))
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
	ctx := r.Context()
	requestId := services.GetRequestId(ctx)
	h.logger.Info("Start method handleLoginUser.", slog.String("request_id", requestId))
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginUser api.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		h.logger.Error(
			"Failed to decode json in handleLoginUser.",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Couldn't decode json body", http.StatusBadRequest)
		return
	}

	if !auth.IsLoginUserValid(loginUser) {
		http.Error(w, "Email and password requiered.", http.StatusBadRequest)
		return
	}

	user, err := h.repository.GetUserByEmail(loginUser.Email)
	if err != nil {
		h.logger.Error(
			"Failed to execute GetUserByEmail in handleLoginUser",
			slog.String("request_id",
				requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Invalid password or Login.", http.StatusUnauthorized)
		return
	}

	if !auth.VerifyPassword(loginUser.Password, user.Password) {
		http.Error(w, "Invalid password or Login.", http.StatusUnauthorized)
		return
	}

	sessionId, expiresAt, err := h.auth.CreateSession(&user)
	if err != nil {
		h.logger.Error(
			"Failed create session for user.",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
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
	http.SetCookie(w, &http.Cookie{
		Name:     auth.SESSION_ID_KEY,
		Value:    sessionId,
		Path:     "/",
		HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	h.logger.Info("Finished method handleLoginUser.", slog.String("request_id", requestId))
}

// handleLogOut godoc
// @Summary      Logout User
// @Tags         users
// @Success      200   "Logged out successfully"
// @Failure      400   {string}  string  "Bad request"
// @Failure      401   {string}  string  "Unauthorized"
// @Failure      500   {string}  string  "Internal server error"
// @Router       /user/info [post]
func (h *Handler) handleLogOut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := services.GetRequestId(ctx)
	h.logger.Info("Start method handleLogOut.", slog.String("request_id", requestId))
	sessionData, err := h.auth.ValidateSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.auth.DeleteSessionById(sessionData.SessionId)
	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:     auth.SESSION_ID_KEY,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
	h.logger.Info("Finished method handleLogOut.", slog.String("request_id", requestId))
}

// handleLoginUser godoc
// @Summary      Ger User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      api.GetUserByIdRequest      true  "User Email"
// @Success      200   {object}  api.User
// @Failure      400   {string}  string  "Bad request"
// @Failure      401   {string}  string  "Unauthorized"
// @Failure      500   {string}  string  "Internal server error"
// @Router       /user/info [post]
func (h *Handler) handleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := services.GetRequestId(ctx)
	h.logger.Info("Start method handleGetUserByEmail.", slog.String("request_id", requestId))

	_, err := h.auth.ValidateSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var getUserByIdRequest api.GetUserByIdRequest
	err = json.NewDecoder(r.Body).Decode(&getUserByIdRequest)
	if err != nil {
		h.logger.Info(
			"Failed to decode json in handleGetUserByEmai.",
			slog.String("request_id", requestId),
			slog.String("Error:", err.Error()))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := h.repository.GetUserByEmail(getUserByIdRequest.Email)
	if err != nil {
		h.logger.Info(
			"Failed to GetUserByEmail in handleGetUserByEmai.",
			slog.String("request_id", requestId),
			slog.String("Error:", err.Error()))
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	h.logger.Info("Finished method handleGetUserByEmail.", slog.String("request_id", requestId))
}


func (h *Handler) handleGetUserProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := services.GetRequestId(ctx)
	sessionData, err := h.auth.ValidateSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.repository.GetUserById(sessionData.UserID)
	if err != nil {
		h.logger.Info(
			"Failed to GetUserById in handleGetUserProfile.",
			slog.String("request_id", requestId),
			slog.String("Error:", err.Error()))
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	userResponse := api.GetUserProfileResponse{
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Email:      user.SecondName,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
	w.WriteHeader(http.StatusOK)
}
