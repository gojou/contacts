package util

// DO NOT DELETE -- handlers utilities

import (
	"html/template"
	"net/http"
)

// Xecute does things
func Xecute(t *template.Template, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t.Execute(w, nil)
	}
}
