package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"shop/internal/api"
	"shop/internal/product"
	"shop/internal/services"
)

func (h *Handler) handleCreateNewProduct(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	requestId := services.GetRequestId(ctx)
	_, err := h.auth.ValidateSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var createProductRequest api.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&createProductRequest); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	product := product.MapCreatProductRequestToProduct(createProductRequest)
	if err := h.repository.CreateProduct(product); err != nil {
		h.logger.Error(
			"Couldn't create producgt",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) handleGetProductsByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := services.GetRequestId(ctx)
	_, err := h.auth.ValidateSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var getProductsRequest api.GetProductsRequest
	if err := json.NewDecoder(r.Body).Decode(&getProductsRequest); err == nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	products, err := h.repository.GetProductsByName(getProductsRequest.Name)
	if err != nil {
		h.logger.Error(
			"Couldn't get product",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	productsResponse := product.MapProductToGetProductsResponse(products)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productsResponse)
	w.WriteHeader(http.StatusOK)
}
