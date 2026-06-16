package handlers

import (
	"fmt"
	"net/http"
	"shop/internal/repository"

	"github.com/go-redis/redis/v8"
)

type Handler struct {
	redis      *redis.Client
	repository *repository.Repository
}

func AddRouter(mux *http.ServeMux, repository *repository.Repository, redis *redis.Client) *http.ServeMux {
	if mux == nil {
		return mux
	}

	handler := newHandler(repository, redis)
	mux.HandleFunc("GET /", handler.handleHome)
	mux.HandleFunc("POST /user/register", handler.handleRegisterUser)
	return mux
}

func newHandler(repository *repository.Repository, redis *redis.Client) *Handler {
	return &Handler{
		redis:      redis,
		repository: repository,
	}
}

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home, page")
}
