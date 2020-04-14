package handlers

import (
	"html/template"
	"net/http"

	"github.com/gojou/contacts/pkg/handlers/util"
)

// Home displays the default "/" page
func Home(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/home.html",
	))

	util.Xecute(page, w, r)

}
