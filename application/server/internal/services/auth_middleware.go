package services

import (
	"net/http"
	"shop/internal/auth"
)

// Maybe I will change it in future
func AuthIsReqiuered(
	authService *auth.Service,
	handler func(http.ResponseWriter, *http.Request,
)) func(w http.ResponseWriter, r *http.Request) {
	if handler == nil {
		panic("No handler is provided")
	}
	if authService == nil {
		panic("No auth service is provided")
	}
	next := http.HandlerFunc(handler)

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := authService.ValidateSession(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
