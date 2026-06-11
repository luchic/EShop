package handlers

import (
	"fmt"
	"net/http"
	"shop/internal/repository"
)

type Handler struct {
	repository *repository.Repository
}

func AddRouter(mux *http.ServeMux, repository *repository.Repository) *http.ServeMux {
	if mux == nil {
		return mux
	}

	handler := newHandler(repository)
	mux.HandleFunc("GET /", handler.handleHome)
	mux.HandleFunc("POST /user/register", handler.handleCreateUser)
	return mux
}

func newHandler(repository *repository.Repository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home, page")
}
