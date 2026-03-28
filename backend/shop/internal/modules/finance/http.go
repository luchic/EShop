package finance

import (
	financeapi "backend/shop/internal/api/finance"
	"backend/shop/internal/repository"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handlers struct {
	repo repository.Repository
}

func NewHandlers(repo repository.Repository) *Handlers {
	return &Handlers{
		repo: repo,
	}
}

func Routers(mux *http.ServeMux, repo repository.Repository) *http.ServeMux {
	if mux == nil {
		return mux
	}

	handlers := NewHandlers(repo)
	mux.HandleFunc("/mony/user/{id}", handlers.handleGetUserBalance)
	return mux
}

func (h *Handlers) handleGetUserBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idValue := r.PathValue("id")
	if idValue == "" {
		http.Error(w, "Missing user id", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(idValue, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	response, ok := h.repo.GetUserBalance(userID)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(financeapi.UserBalanceResponse{
		UserID:  response.UserID,
		Balance: response.Balance,
	}); err != nil {
		http.Error(w, "Couldn't encode data", http.StatusInternalServerError)
		return
	}
}
