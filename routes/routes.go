package routes

import (
	"github.com/go-chi/chi"

	"github.com/CristianSch248/App_GoLang/handlers"
)

func GetRoutes() *chi.Mux {
	route := chi.NewRouter()

	route.Get("/musicas", handlers.GetMusicas)

	return route
}
