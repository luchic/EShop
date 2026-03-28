package finance

import (
	"backend/shop/internal/repository"
	"net/http"
)

type Handlers struct {
	repo repository.Repository
}

func NewHandlers(repo repository.Repository) *Handlers {
	return &Handlers{
		repo: repo,
	}
}

func Routers(mux *http.ServeMux, repo repository.Repository) *http.ServeMux {
	if mux == nil {
		return mux
	}

	handlers := NewHandlers(repo)
	_ = handlers
	return mux
}
