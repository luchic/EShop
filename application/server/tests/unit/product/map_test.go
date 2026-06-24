package product

import (
	"shop/internal/api"
	"shop/internal/product"
	"testing"
)

func TestMapCreateProductRequestToProduct(t *testing.T) {
	request := api.CreateProductRequest{
		Name:        "Test Product",
		Description: "A test product",
		Price:       19.99,
		Stock:       100,
	}

	result := product.MapCreatProductRequestToProduct(request)

	if result.Name != request.Name {
		t.Errorf("Name = %q, want %q", result.Name, request.Name)
	}
	if result.Description != request.Description {
		t.Errorf("Description = %q, want %q", result.Description, request.Description)
	}
	if result.Price != request.Price {
		t.Errorf("Price = %v, want %v", result.Price, request.Price)
	}
	if result.Stock != request.Stock {
		t.Errorf("Stock = %v, want %v", result.Stock, request.Stock)
	}
	if result.ImageUrl != "" {
		t.Errorf("ImageUrl = %q, want empty string", result.ImageUrl)
	}
}

func TestMapProductToGetProductsResponse(t *testing.T) {
	p := api.Product{
		Id:          1,
		Name:        "Test Product",
		Description: "A test product",
		Price:       29.99,
		Stock:       50,
		ImageUrl:    "http://example.com/image.png",
	}

	result := product.MapProductToGetProductsResponse(p)

	if result.Id != p.Id {
		t.Errorf("Id = %v, want %v", result.Id, p.Id)
	}
	if result.Name != p.Name {
		t.Errorf("Name = %q, want %q", result.Name, p.Name)
	}
	if result.Description != p.Description {
		t.Errorf("Description = %q, want %q", result.Description, p.Description)
	}
	if result.Price != p.Price {
		t.Errorf("Price = %v, want %v", result.Price, p.Price)
	}
	if result.Stock != p.Stock {
		t.Errorf("Stock = %v, want %v", result.Stock, p.Stock)
	}
}

func TestMapProductArrayToGetProductsResponse_Empty(t *testing.T) {
	var products []api.Product

	result := product.MapProductArrayToGetProductsResponse(products)

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %d items", len(result))
	}
}

func TestMapProductArrayToGetProductsResponse_Multiple(t *testing.T) {
	products := []api.Product{
		{Id: 1, Name: "Product A", Description: "Desc A", Price: 10.0, Stock: 5},
		{Id: 2, Name: "Product B", Description: "Desc B", Price: 20.0, Stock: 10},
		{Id: 3, Name: "Product C", Description: "Desc C", Price: 30.0, Stock: 15},
	}

	result := product.MapProductArrayToGetProductsResponse(products)

	if len(result) != 3 {
		t.Fatalf("expected 3 items, got %d", len(result))
	}

	for i, p := range products {
		if result[i].Id != p.Id {
			t.Errorf("item %d: Id = %v, want %v", i, result[i].Id, p.Id)
		}
		if result[i].Name != p.Name {
			t.Errorf("item %d: Name = %q, want %q", i, result[i].Name, p.Name)
		}
		if result[i].Price != p.Price {
			t.Errorf("item %d: Price = %v, want %v", i, result[i].Price, p.Price)
		}
	}
}
