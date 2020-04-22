package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

// FormData stores field information in the case of a POST failure
type FormData struct {
	ID         string
	FirstName  string
	LastName   string
	BirthYear  string
	BirthMonth string
	BirthDay   string
	Email      string
}

// Data for use in the form
var data FormData

//Contact displays the contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering the Contact form %s", data.ID)
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact.html",
	))

	page.Execute(w, data)

}

// ContactPost handles a POST from the Contact form
func ContactPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for key, value := range r.Form {
		if value[0] != "" {
			log.Printf("key = %s -- value = %s", key, value[0])
		}
	}

	data.ID = r.FormValue("id")
	log.Printf("Final test, structure formdata.id = %s", data.ID)
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact.html",
	))

	// validate returns a slice of errors if the form contains any invalid
	// strings.
	// TODO: Update contact.html to populate the VALID data from the original
	// form submission.
	if len(validate(data)) > 0 {
		page.Execute(w, data)
	} else {
		// TODO: save the form data to the database
		// e:= repository.Commit(data)
		// if e!=nil {} etc
		data = FormData{}
		page.Execute(w, data)
	}
}

func validate(data FormData) (errs []error) {
	return append(errs, errors.New("danger will robinson"))
}
