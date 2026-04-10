package main

import (
	"backend/shop/internal/config"
	"backend/shop/internal/modules/auth"
	"backend/shop/internal/modules/finance"
	"backend/shop/internal/modules/goods"
	"backend/shop/internal/modules/swagger"
	"backend/shop/internal/repository"
	"context"
	"log/slog"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Couldn't load config", slog.Any("err", err))
		return
	}
	ctx := context.Background()

	repo, err := repository.NewPostgresRepository(ctx, cfg)
	if err != nil {
		slog.Error("Couldn't load repository", slog.Any("err", err))
		return
	}

	mux := http.NewServeMux()
	mux = swagger.Routers(mux)
	mux = auth.Routers(mux, repo, cfg)
	mux = goods.Routers(mux, repo)
	mux = finance.Routers(mux, repo)

	slog.Info("Run server...\n")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error("Couldn't run server", slog.Any("err", err))
	}
}
