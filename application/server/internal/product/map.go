package product

import "shop/internal/api"

func MapCreatProductRequestToProduct(requet api.CreateProductRequest) api.Product {
	return api.Product{
		Name:        requet.Name,
		Description: requet.Description,
		Price:       requet.Price,
		Stock:       requet.Stock,
		ImageUrl:    "",
	}
}
