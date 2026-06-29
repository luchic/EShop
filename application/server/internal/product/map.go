package product

import (
	"database/sql"
	"shop/internal/api"
)

func MapCreateProductRequestToProduct(requet api.CreateProductRequest) api.Product {
	return api.Product{
		Name:        requet.Name,
		Description: sql.NullString{String: requet.Description, Valid: true},
		Price:       requet.Price,
		Stock:       requet.Stock,
		ImageUrl:    sql.NullString{Valid: false},
	}
}

func MapProductArrayToGetProductsResponse(products []api.Product) []api.GetProductsResponse {
	productsResponse := make([]api.GetProductsResponse, len(products))
	for i := 0; i < len(products); i++ {
		productsResponse[i] = MapProductToGetProductsResponse(products[i])
	}
	return productsResponse
}

func MapProductToGetProductsResponse(product api.Product) api.GetProductsResponse {
	var description string = ""
	if product.Description.Valid {
		description = product.Description.String
	}

	return api.GetProductsResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
