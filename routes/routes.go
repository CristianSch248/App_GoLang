package routes

import (
	"github.com/go-chi/chi"

	"github.com/CristianSch248/App_GoLang/handlers"
)

func GetRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/hello", handlers.HelloHandler)
	r.Get("/another", handlers.AnotherHandler)

	return r
}
