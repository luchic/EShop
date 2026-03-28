package repository

import "backend/shop/internal/api/goods"

type MemoryRepository struct {
	items []goods.Product
}

func NewMemoryRepository(items []goods.Product) *MemoryRepository {
	return &MemoryRepository{items: items}
}

func (r *MemoryRepository) AddProduct(product goods.AddProductRequest) {
	lastId := r.items[len(r.items)-1].Id + 1
	repositoryProduct := goods.Product{
		Id: lastId,
		Name: product.Name,
		Description: product.Description,
	}
	r.items = append(r.items, repositoryProduct)
}	

func (r *MemoryRepository) GetGoodPage(offset int, limit int) []goods.Product {
	if offset >= len(r.items) {
		return []goods.Product{}
	}

	end := offset + limit
	if end > len(r.items) {
		end = len(r.items)
	}

	return r.items[offset:end]
}
