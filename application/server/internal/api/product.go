package api

import "time"

type Product struct {
	Id          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float32   `db:"price"`
	Stock       int32     `db:"stock"`
	ImageUrl    *string   `db:"image_url"`
	CreatedAt   time.Time `db:"created_at"`
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int32   `json:"stock"`
}

type GetProductsRequest struct {
	Name string `json:"name"`
}

type GetProductsResponse struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int32   `json:"stock"`
}
