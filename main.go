package main

import (
	"fmt"
	"net/http"

	"github.com/CristianSch248/App_GoLang/configs"
	"github.com/CristianSch248/App_GoLang/routes"
	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Mount("/", routes.GetRoutes())

	serverAddr := fmt.Sprintf(":%s", configs.GetServerPort())
	fmt.Printf("Servidor escutando em http://localhost%s\n", serverAddr)

	err = http.ListenAndServe(serverAddr, r)
	if err != nil {
		panic(err)
	}
}
