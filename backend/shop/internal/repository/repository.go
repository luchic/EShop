package repository

import (
	authapi "backend/shop/internal/api/auth"
	financeapi "backend/shop/internal/api/finance"
	"backend/shop/internal/api/goods"
)

type Repository interface {
	GetProductPage(offset int, limit int) []goods.Product
	GetProductByID(id uint64) (goods.Product, bool)
	AddProduct(product goods.AddProductRequest) (goods.AddProductResponse, error)
	GetUserBalance(userID int64) (financeapi.UserBalanceResponse, bool)
	RegisterTransaction(request financeapi.RegisterTransactionRequest) (financeapi.Transaction, error)
	GetTransactionByID(id int64) (financeapi.Transaction, bool)
	GetTransactionsByUserID(userID int64) ([]financeapi.Transaction, error)
	SaveOAuthState(state string) error
	HasOAuthState(state string) bool
	DeleteOAuthState(state string) error
	UpsertOAuthUser(user authapi.OAuthUser) (authapi.AppUser, error)
}
