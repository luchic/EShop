package goods

import (
	"backend/shop/internal/api/goods"
	"backend/shop/internal/repository"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type Handlers struct {
	repo repository.Repository
}

func Routers(repo repository.Repository) *http.ServeMux {
	handlers := NewHandlers(repo)
	mux := http.NewServeMux()
	mux.HandleFunc("/itmes", handlers.handleGetAllProducts)
	mux.HandleFunc("/itmes/add", handlers.handleAddProduct)
	return mux
}

func NewHandlers(repo repository.Repository) *Handlers {
	return &Handlers{
		repo: repo,
	}
}

func getIntQuery(r *http.Request, key string) (int, error) {
	query := r.URL.Query()
	value := query.Get(key)
	if value == "" {
		return 0, fmt.Errorf("No value")
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func (h *Handlers) handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	slog.Info("Hanele getting all products\n")
	page := 1
	limit := 4

	value, err := getIntQuery(r, "page")
	if err == nil && value > 1 {
		page = value
	}
	value, err = getIntQuery(r, "limit")
	if err == nil && value > 1 {
		limit = value
	}
	offset := (page - 1) * limit
	products := h.repo.GetGoodPage(offset, limit)

	pageResponse := goods.PaginatedResponse{
		Items: products,
		Page:  page,
		Limit: limit,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(pageResponse)
	if err != nil {
		slog.Error("Failed to encode products\n")
		http.Error(w, "Couldn't endcode data", http.StatusBadRequest)
		return
	}
	slog.Info("Hanele getting all products operation is finished\n")
}

func (h *Handlers) handleAddProduct(w http.ResponseWriter, r *http.Request) {
	slog.Info("handle add product operation\n")

	if r.Method != http.MethodPost {
		slog.Error("handle add product operation: Method not allowed\n")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request goods.AddProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("handle add product operation: Couldn't decode json body\n")
		http.Error(w, "Couldn't decode json body", http.StatusBadRequest)
		return
	}
	h.repo.AddProduct(request)
	slog.Info("handle add product operation is finished\n")
}
