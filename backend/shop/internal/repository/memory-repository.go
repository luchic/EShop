package repository

import (
	authapi "backend/shop/internal/api/auth"
	"backend/shop/internal/api/goods"
)

type MemoryRepository struct {
	items       []goods.Product
	oauthStates map[string]struct{}
	users       map[int64]authapi.AppUser
}

func NewMemoryRepository(items []goods.Product) *MemoryRepository {
	return &MemoryRepository{
		items:       items,
		oauthStates: map[string]struct{}{},
		users:       map[int64]authapi.AppUser{},
	}
}

func (r *MemoryRepository) AddProduct(product goods.AddProductRequest) {
	lastId := r.items[len(r.items)-1].Id + 1
	repositoryProduct := goods.Product{
		Id:          lastId,
		Name:        product.Name,
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

func (r *MemoryRepository) GetGoodByID(id uint64) (goods.Product, bool) {
	for _, product := range r.items {
		if product.Id == id {
			return product, true
		}
	}

	return goods.Product{}, false
}

func (r *MemoryRepository) SaveOAuthState(state string) error {
	r.oauthStates[state] = struct{}{}
	return nil
}

func (r *MemoryRepository) HasOAuthState(state string) bool {
	_, ok := r.oauthStates[state]
	return ok
}

func (r *MemoryRepository) DeleteOAuthState(state string) error {
	delete(r.oauthStates, state)
	return nil
}

func (r *MemoryRepository) UpsertOAuthUser(user authapi.OAuthUser) (authapi.AppUser, error) {
	appUser := authapi.AppUser{
		ID:          user.ProviderID,
		Login:       user.Login,
		DisplayName: user.DisplayName,
		Email:       user.Email,
	}
	r.users[user.ProviderID] = appUser
	return appUser, nil
}
