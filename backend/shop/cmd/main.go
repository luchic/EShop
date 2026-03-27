package cmd

import (
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

	mux := goods.Routers()
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Couldn't run server: %v", err)
	}
}
