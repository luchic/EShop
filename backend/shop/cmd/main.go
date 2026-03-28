package main

import (
	apiGood "backend/shop/internal/api/goods"
	"backend/shop/internal/modules/goods"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// cfg, err := config.Load()
	// if err != nil {
	// 	log.Printf("load config: %v", err)
	// }
	// _ = cfg

	repo := goods.NewMemoryRepository([]apiGood.Product{
		{Id: 1, Name: "Phone", Description: "Smartphone"},
		{Id: 2, Name: "Keyboard", Description: "Mechanical keyboard"},
		{Id: 3, Name: "Mouse", Description: "Wireless mouse"},
		{Id: 4, Name: "Monitor", Description: "27 inch monitor"},
		{Id: 5, Name: "Printer", Description: "27 inch monitor"},
		{Id: 6, Name: "Laptop", Description: "27 inch monitor"},
		{Id: 7, Name: "PC", Description: "27 inch monitor"},
	})

	mux := goods.Routers(repo)
	fmt.Printf("Run server...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Couldn't run server: %v", err)
	}
}
