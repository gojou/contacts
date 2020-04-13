package handlers

import (
	"html/template"
	"net/http"
)

//Contact displays the contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact.html",
	))

	xecute(page, w, r)

}
