package handlers

import (
	"html/template"
	"net/http"
)

//About displays the "About" page
func About(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/about.html",
	))

	xecute(page, w, r)

}
