package router

import (
	"net/http"

	"github.com/Grizak/go-testing/internal/handlers"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	r.Get("/", handlers.Home)
	r.Get("/about", handlers.About)

	// Serve static files
	fs := http.FileServer(http.Dir("./web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	return r
}
