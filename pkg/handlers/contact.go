package handlers

import (
	"html/template"
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
