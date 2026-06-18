package services

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

const contextKeyRequestId string = "RequestId"

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := uuid.New()

		ctx = context.WithValue(ctx, contextKeyRequestId, id.String())

		r = r.WithContext(ctx)

		slog.Debug(
			"Incomming request",
			slog.String("Method", r.Method),
			slog.String("Request URI", r.RequestURI),
			slog.String("Request URI", r.RemoteAddr),
			slog.String("Request Id", id.String()))

		next.ServeHTTP(w, r)

		slog.Debug(
			"Finished handling http req.",
			slog.String("Request Id", id.String()))
	})
}

func GetRequestId(ctx context.Context) string {
	requestId := ctx.Value(contextKeyRequestId)

	value, ok := requestId.(string)
	if ok {
		return value
	}
	return ""
}
