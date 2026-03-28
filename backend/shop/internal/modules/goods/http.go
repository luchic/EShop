package goods

import (
	"backend/shop/internal/api/goods"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handlers struct {
	repo Repository
}

func Routers(repo Repository) *http.ServeMux {
	handlers := NewHandlers(repo)
	mux := http.NewServeMux()
	mux.HandleFunc("/itmes", handlers.handleGetAllProducts)
	mux.HandleFunc("/itmes/add", handlers.handleAddProduct)
	return mux
}

func NewHandlers(repo Repository) *Handlers {
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

	page := 1
	limit := 4

	value, err := getIntQuery(r, "page")
	if err != nil && value > 1 {
		page = value
	}
	value, err = getIntQuery(r, "limit")
	if err != nil && value > 1 {
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
		http.Error(w, "Couldn't endcode data", http.StatusBadRequest)
		return
	}
}

func (h *Handlers) handleAddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
