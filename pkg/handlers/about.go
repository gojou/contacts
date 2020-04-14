package handlers

import (
	"html/template"
	"net/http"

	"github.com/gojou/contacts/pkg/handlers/util"
)

//About displays the "About" page
func About(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/about.html",
	))

	util.Xecute(page, w, r)

}
