package handlers

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func AnotherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is another route"))
}
