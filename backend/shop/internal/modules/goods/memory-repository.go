package goods

import "backend/shop/internal/api/goods"

type MemoryRepository struct {
	items []goods.Product
}

func NewMemoryRepository(items []goods.Product) *MemoryRepository {
	return &MemoryRepository{items: items}
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
