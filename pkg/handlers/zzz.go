package handlers

import (
	"html/template"
	"net/http"
)

func xecute(t *template.Template, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t.Execute(w, nil)
	}
}
