package handlers

import (
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
}

func (handler *Handler) AddRouter(mux *http.ServeMux) *http.ServeMux {
	if mux == nil {
		return mux
	}

	mux.HandleFunc("GET /", handler.handleHome)
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)
	return mux
}

func NewDefualtHandler() *Handler {
	return &Handler{}
}

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home, page")
}
