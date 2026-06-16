// @title           Shop API
// @version         1.0
// @description     Theodor's Emporium backend API.
//
// @host      localhost:8080
// @BasePath  /

package main

import (
	"fmt"
	"net/http"
	"shop/internal/config"
	"shop/internal/handlers"
	"shop/internal/repository"
	_ "shop/docs"
)

func main() {
	cfg, err := config.NewConfig("config.json")
	if err != err {
		fmt.Println("Couldn't create configuration structure")
		return
	}

	repo, err := repository.NewRepository(cfg)
	if err != nil {
		fmt.Println("Couldn't create connection to database: ", err)
		return
	}

	defer repo.Close()

	redis, err := repository.NewRedis(cfg)
	if err != nil {
		fmt.Println("Couldn't connect to redis:", err)
		return
	}

	defer redis.Close()

	mux := http.NewServeMux()
	handlers.AddRouter(mux, repo, redis)

	fmt.Printf("Listen %s\n", cfg.Host)
	http.ListenAndServe(cfg.Host, mux)
}
