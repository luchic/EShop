package finance

import (
	financeapi "backend/shop/internal/api/finance"
	"backend/shop/internal/repository"
	"encoding/json"
	"errors"
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
	mux.HandleFunc("/mony/user", handlers.handleUpdateUserBalance)
	mux.HandleFunc("/mony/user/{id}", handlers.handleGetUserBalance)
	mux.HandleFunc("/transaction/register", handlers.handleRegisterTransaction)
	mux.HandleFunc("/transaction/user/{id}", handlers.handleGetTransactionsByUserID)
	mux.HandleFunc("/transaction/{id}", handlers.handleGetTransactionByID)
	return mux
}

func (h *Handlers) handleUpdateUserBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request financeapi.UpdateUserBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, "Couldn't decode json body", http.StatusBadRequest)
		return
	}

	response, err := h.repo.UpdateUserBalance(request)
	if err != nil {
		writeRepositoryError(w, err, "Couldn't update user balance")
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *Handlers) handleGetUserBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idValue := r.PathValue("id")
	if idValue == "" {
		writeError(w, "Missing user id", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(idValue, 10, 64)
	if err != nil {
		writeError(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	response, ok := h.repo.GetUserBalance(userID)
	if !ok {
		writeError(w, "User not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, financeapi.UserBalanceResponse{
		UserID:  response.UserID,
		Balance: response.Balance,
	})
}

func (h *Handlers) handleRegisterTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request financeapi.RegisterTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, "Couldn't decode json body", http.StatusBadRequest)
		return
	}

	transaction, err := h.repo.RegisterTransaction(request)
	if err != nil {
		writeRepositoryError(w, err, "Couldn't register transaction")
		return
	}

	writeJSON(w, http.StatusCreated, transaction)
}

func (h *Handlers) handleGetTransactionByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idValue := r.PathValue("id")
	if idValue == "" {
		writeError(w, "Missing transaction id", http.StatusBadRequest)
		return
	}

	transactionID, err := strconv.ParseInt(idValue, 10, 64)
	if err != nil {
		writeError(w, "Invalid transaction id", http.StatusBadRequest)
		return
	}

	transaction, ok := h.repo.GetTransactionByID(transactionID)
	if !ok {
		writeError(w, "Transaction not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, transaction)
}

func (h *Handlers) handleGetTransactionsByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idValue := r.PathValue("id")
	if idValue == "" {
		writeError(w, "Missing user id", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(idValue, 10, 64)
	if err != nil {
		writeError(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	transactions, err := h.repo.GetTransactionsByUserID(userID)
	if err != nil {
		writeRepositoryError(w, err, "Couldn't fetch transactions")
		return
	}

	writeJSON(w, http.StatusOK, transactions)
}

func writeRepositoryError(w http.ResponseWriter, err error, fallbackMessage string) {
	switch {
	case errors.Is(err, repository.ErrUserNotFound):
		writeError(w, "User not found", http.StatusNotFound)
	case errors.Is(err, repository.ErrProductNotFound):
		writeError(w, "Product not found", http.StatusNotFound)
	case errors.Is(err, repository.ErrInsufficientFunds):
		writeError(w, "Insufficient funds", http.StatusBadRequest)
	case errors.Is(err, repository.ErrPriceMismatch):
		writeError(w, "Price mismatch", http.StatusBadRequest)
	default:
		writeError(w, fallbackMessage, http.StatusInternalServerError)
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Couldn't encode data", http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, message string, status int) {
	http.Error(w, message, status)
}
