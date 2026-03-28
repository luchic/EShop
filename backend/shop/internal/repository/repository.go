package repository

import (
	authapi "backend/shop/internal/api/auth"
	"backend/shop/internal/api/goods"
)

type Repository interface {
	GetGoodPage(offset int, limit int) []goods.Product
	GetGoodByID(id uint64) (goods.Product, bool)
	AddProduct(product goods.AddProductRequest)
	SaveOAuthState(state string) error
	HasOAuthState(state string) bool
	DeleteOAuthState(state string) error
	UpsertOAuthUser(user authapi.OAuthUser) (authapi.AppUser, error)
}
