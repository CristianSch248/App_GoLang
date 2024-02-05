package handlers

import (
	"fmt"
	"net/http"

	"github.com/CristianSch248/App_GoLang/db"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConn()

	if err != nil {
		// Lide com o erro de maneira apropriada, como enviar uma resposta de erro ao cliente
		http.Error(w, "Erro na conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	fmt.Println("aqui?", db)

	w.Write([]byte("Hello, World!"))
}

func AnotherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is another route"))
}
