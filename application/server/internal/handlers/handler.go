package handlers

import (
	"fmt"
	"net/http"
	"shop/internal/auth"
	"shop/internal/repository"

	"github.com/go-redis/redis/v8"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	auth       *auth.Service
	repository *repository.Repository
}

func AddRouter(mux *http.ServeMux, repository *repository.Repository, redis *redis.Client) *http.ServeMux {
	if mux == nil {
		return mux
	}

	handler := newHandler(repository, redis)
	mux.HandleFunc("GET /", handler.handleHome)
	mux.HandleFunc("POST /user/register", handler.handleRegisterUser)
	mux.HandleFunc("POST /user/login", handler.handleLoginUser)
	mux.HandleFunc("POST /user/info", handler.handleGetUserByEmail)
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)
	return mux
}

func newHandler(repository *repository.Repository, redis *redis.Client) *Handler {
	return &Handler{
		auth:       auth.NewService(redis),
		repository: repository,
	}
}

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home, page")
}
