package main

import (
	"net/http"

	"github.com/gojou/contacts/pkg/handlers"
	"github.com/gorilla/mux"
)

func routing(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/contact", handlers.Contact).Methods("GET")
	r.HandleFunc("/contact", handlers.ContactPost).Methods("POST")
	r.HandleFunc("/about", handlers.About)
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

}
