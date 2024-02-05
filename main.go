package main

import (
	"net/http"

	"github.com/CristianSch248/App_GoLang/routes"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Mount("/", routes.GetRoutes())

	http.ListenAndServe(":8080", r)
}
