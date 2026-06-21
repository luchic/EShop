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

func MapProductArrayToGetProductsResponse(products []api.Product) []api.GetProductsResponse {
	productsResponse := make([]api.GetProductsResponse, len(products))
	for i := 0; i < len(products); i++ {
		productsResponse[i] = MapProductToGetProductsResponse(products[i])
	}
	return productsResponse
}

func MapProductToGetProductsResponse(product api.Product) api.GetProductsResponse {
	return api.GetProductsResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
