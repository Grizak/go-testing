package main

import (
	"log"
	"net/http"

	"github.com/Grizak/go-testing/internal/config"
	"github.com/Grizak/go-testing/internal/router"
)

func main() {
	// Load env config
	cfg := config.LoadConfig()

	// Create router
	r := router.NewRouter()

	log.Printf("Server listening on %s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
