package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gojou/contacts/pkg/handlers/util"
)

//Contact displays the contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact.html",
	))

	util.Xecute(page, w, r)

}

// ContactPost handles a POST from the Contact form
func ContactPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for key, value := range r.Form {
		if value[0] != "" {
			log.Printf("key = %s -- value = %s", key, value)
		}
	}

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact.html",
	))

	page.Execute(w, nil)

}
