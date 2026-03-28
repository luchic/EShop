package repository

import "backend/shop/internal/api/goods"

type Repository interface {
	GetGoodPage(offset int, limit int) []goods.Product
	AddProduct(product goods.AddProductRequest)
}
