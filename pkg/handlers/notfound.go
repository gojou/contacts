package handlers

import (
	"html/template"
	"net/http"

	"github.com/gojou/contacts/pkg/handlers/util"
)

// NotFound handles 404
func NotFound(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/notfound.html",
	))
	w.WriteHeader(http.StatusNotFound)
	util.Xecute(page, w, r)
}
