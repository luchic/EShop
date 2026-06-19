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

func MapProductToGetProductsResponse(products []api.Product) []api.GetProductsResponse {
	productsResponse := make([]api.GetProductsResponse, len(products))
	for i := 0; i < len(products); i++ {
		productsResponse[i] = api.GetProductsResponse{
			Id:          products[i].Id,
			Name:        products[i].Name,
			Description: products[i].Description,
			Price:       products[i].Price,
			Stock:       products[i].Stock,
		}
	}
	return productsResponse
}
