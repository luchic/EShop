// @title           Shop API
// @version         1.0
// @description     Theodor's Emporium backend API.
//
// @host      localhost:8080
// @BasePath  /

package main

import (
	"fmt"
	"log/slog"
	"net/http"
	_ "shop/docs"
	"shop/internal/auth"
	"shop/internal/config"
	"shop/internal/handlers"
	"shop/internal/repository"
	"shop/internal/services"
)

func main() {
	cfg, err := config.NewConfig("config.json")
	if err != nil {
		fmt.Println("Couldn't create configuration structure")
		return
	}

	db, err := repository.NewDB(cfg)
	if err != nil {
		fmt.Println("Couldn't create connection to database: ", err)
		return
	}
	defer db.Close()

	redis, err := repository.NewRedis(cfg)
	if err != nil {
		fmt.Println("Couldn't connect to redis:", err)
		return
	}

	defer redis.Close()

	logger := slog.Default()
	authService := auth.NewService(redis)

	userRepo := repository.NewPostgresUserRepository(db)
	productRepo := repository.NewPostgresProductRepository(db)

	DefaultHandler := handlers.NewDefaultHandler()
	userHandler := handlers.NewUserHandler(logger, authService, userRepo)
	productHandler := handlers.NewProductHandler(logger, authService, productRepo)

	mux := http.NewServeMux()
	DefaultHandler.AddRouter(mux)
	userHandler.AddUserHandlerRouter(mux)
	productHandler.AddProductHandlerRouter(mux)

	wrappedMux := services.RequestIdMiddleware(mux)
	fmt.Printf("Listen %s\n", cfg.Host)
	http.ListenAndServe(cfg.Host, wrappedMux)
}
