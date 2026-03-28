package cmd

import (
	apiGood "backend/shop/internal/api/goods"
	"backend/shop/internal/config"
	"backend/shop/internal/modules/goods"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("load config: %v", err)
	}
	_ = cfg

	repo := goods.NewMemoryRepository([]apiGood.Product{
		{Id: 1, Name: "Phone", Description: "Smartphone"},
		{Id: 2, Name: "Keyboard", Description: "Mechanical keyboard"},
		{Id: 3, Name: "Mouse", Description: "Wireless mouse"},
		{Id: 4, Name: "Monitor", Description: "27 inch monitor"},
	})

	mux := goods.Routers(repo)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Couldn't run server: %v", err)
	}
}
