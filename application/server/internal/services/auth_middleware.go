package services

import (
	"context"
	"fmt"
	"net/http"
	"shop/internal/api"
	"shop/internal/auth"
)

const sessionKey string = "session"

// Maybe I will change it in future
func AuthIsRequired(
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
		session, err := authService.ValidateSession(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, sessionKey, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func GetSessionDataFromContext(ctx context.Context) (api.SessionData, error) {
	session := ctx.Value(sessionKey)

	value, ok := session.(api.SessionData)
	if !ok {
		return api.SessionData{}, fmt.Errorf("No session data in context")
	}
	return value, nil
}
