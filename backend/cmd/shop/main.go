package shop

import (
	"fmt"
	"net/http"
	"shop/internal/config"
	"shop/internal/handlers"
	"shop/internal/repository"
)

func main() {
	cfg, err := config.NewConfig("config.json")
	if err != err {
		fmt.Println("Couldn't create configuration structure")
		return
	}

	repository, err := repository.NewRepository(cfg)
	if err != nil {
		fmt.Print("Couldn't create connection to database")
		return
	}

	defer repository.Close()

	mux := http.NewServeMux()
	handlers.AddRouter(mux, repository)

	const addr = ":8080"
	http.ListenAndServe(addr, mux)
}
