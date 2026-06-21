package repository

import "shop/internal/api"

type UserRepository interface {
	CreateUser(user api.User) error
	GetUserByEmail(email string) (api.User, error)
	GetUserById(id int64) (api.User, error)
}

type ProductRepository interface {
	CreateProduct(product api.Product) error
	GetProductById(id int64) (api.Product, error)
	GetProductsByName(name string) ([]api.Product, error)
}
