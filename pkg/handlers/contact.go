package handlers

import (
	"html/template"
	"log"
	"net/http"

	val9 "gopkg.in/go-playground/validator.v9"
)

// FormData stores field information in the case of a POST failure
type FormData struct {
	ID         string `validate:"required"`
	FirstName  string `validate:"required"`
	LastName   string `validate:"required"`
	BirthYear  string `validate:"omitempty,len=4"`
	BirthMonth string `validate:"omitempty,len=2"`
	BirthDay   string `validate:"omitempty,len=2"`
	Email      string `json:"email" validate:"required"`
}

// Data for use from the form
var data FormData

//Contact displays the contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering the Contact form %s", data.ID)
	executePage(w, data)

}

// ContactPost handles a POST from the Contact form
func ContactPost(w http.ResponseWriter, r *http.Request) {
	v := val9.New()
	r.ParseForm()
	c := FormData{}

	for key, value := range r.Form {
		if value[0] != "" {
			log.Printf("key = %s -- value = %s", key, value[0])
		}
	}
	c.ID = r.Form.Get("id")
	c.FirstName = r.Form.Get("firstname")
	c.LastName = r.Form.Get("lastname")
	c.BirthYear = r.Form.Get("birthyear")
	c.BirthMonth = r.Form.Get("birthmonth")
	c.BirthDay = r.Form.Get("birthday")
	c.Email = r.Form.Get("email")

	verrs := v.Struct(c)
	if verrs != nil {
		for _, e := range verrs.(val9.ValidationErrors) {
			log.Printf("Validator Error: %s", e)
		}
		data.ID = r.FormValue("id")
		data.FirstName = r.FormValue("firstname")
		data.LastName = r.FormValue("lastname")
		data.BirthYear = r.FormValue("birthyear")
		data.BirthMonth = r.FormValue("birthmonth")
		data.BirthDay = r.FormValue("birthday")
		data.Email = r.FormValue("email")

		executePage(w, data)
	} else {
		data = FormData{}

		executePage(w, data)
	}
}

func executePage(w http.ResponseWriter, data FormData) {
	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/contact.html",
	))

	page.Execute(w, data)

}
