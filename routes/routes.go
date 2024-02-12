package routes

import (
	"github.com/go-chi/chi"

	"github.com/CristianSch248/App_GoLang/handlers"
)

func GetRoutes() *chi.Mux {
	route := chi.NewRouter()

	route.Get("/musicas", handlers.GetMusicas)

	// r.Get("/hello", handlers.HelloHandler)
	// r.Get("/another", handlers.AnotherHandler)

	return route
}
