package main

import (
	"net/http"

	"github.com/gojou/bones/pkg/handlers"
	"github.com/gorilla/mux"
)

func routing(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/contact", handlers.Contact)
	r.HandleFunc("/about", handlers.About)
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

}
