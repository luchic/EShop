package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"shop/internal/api"
	"shop/internal/auth"
	"shop/internal/product"
	"shop/internal/repository"
	"shop/internal/services"
	"strconv"
)

type ProductHandler struct {
	logger     *slog.Logger
	auth       *auth.Service
	repository repository.ProductRepository
}

func NewProductHandler(
	logger *slog.Logger,
	authService *auth.Service,
	productRepo repository.ProductRepository,
) *ProductHandler {
	return &ProductHandler{
		logger:     logger,
		auth:       authService,
		repository: productRepo,
	}
}

func (handler *ProductHandler) AddProductHandlerRouter(mux *http.ServeMux) *http.ServeMux {
	if mux == nil {
		return mux
	}

	mux.HandleFunc("POST /products/create", handler.handleCreateNewProduct)
	mux.HandleFunc("POST /products", handler.handleCreateNewProduct)
	mux.HandleFunc("GET /products/{id}", handler.handleGetProductById)
	mux.HandleFunc("POST /products/search", handler.handleGetProductsByName)

	return mux
}

// handleCreateNewProduct godoc
// @Summary Create a new product
// @Description Creates a new product in the shop catalog
// @Tags products
// @Accept json
// @Produce json
// @Param request body api.CreateProductRequest true "Product to create"
// @Success 201 "Product created"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal Error"
// @Security ApiKeyAuth
// @Router /products [post]
func (h *ProductHandler) handleCreateNewProduct(
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

// handleGetProductsByName godoc
// @Summary Get products by name
// @Description Searches for products matching the given name
// @Tags products
// @Accept json
// @Produce json
// @Param request body api.GetProductsRequest true "Product name to search"
// @Success 200 {array} api.GetProductsResponse "List of matching products"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal Error"
// @Security ApiKeyAuth
// @Router /products/search [post]
func (h *ProductHandler) handleGetProductsByName(w http.ResponseWriter, r *http.Request) {
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
	if err := json.NewDecoder(r.Body).Decode(&getProductsRequest); err != nil {
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
	productsResponse := product.MapProductArrayToGetProductsResponse(products)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productsResponse)
	w.WriteHeader(http.StatusOK)
}

// handleGetProductById godoc
// @Summary Get products by id
// @Description Return the product by id
// @Tags products
// @Produce json
// @Param id path int64 true "Product ID"
// @Success 200 {object} api.GetProductsResponse "Product found"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal Error"
// @Security ApiKeyAuth
// @Router /products/{id} [get]
func (h *ProductHandler) handleGetProductById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := services.GetRequestId(ctx)
	_, err := h.auth.ValidateSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := getIdFromQuery(r)
	if err != nil {
		h.logger.Error(
			"Couldn't parse id from query",
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	productModel, err := h.repository.GetProductById(id)
	if err != nil {
		h.logger.Error(
			"Couldn't get product from database",
			slog.Int64("Product_id", id),
			slog.String("request_id", requestId),
			slog.String("Error", err.Error()))
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	responseProduct := product.MapProductToGetProductsResponse(productModel)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProduct)
	w.WriteHeader(http.StatusOK)
}

func getIdFromQuery(r *http.Request) (int64, error) {
	value := r.PathValue("id")
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1, err
	}
	return num, nil
}
