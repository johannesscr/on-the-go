package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", parseQueryParams)
	http.HandleFunc("/form", form)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type person struct {
	FirstName string
	LastName string
	Subscribed bool
}

func form(w http.ResponseWriter, r * http.Request) {
	f := r.FormValue("firstName")
	l := r.FormValue("lastName")
	s := r.FormValue("subscribed") == "on"

	personData := person{
		FirstName: f,
		LastName: l,
		Subscribed: s,
	}

	err := tpl.ExecuteTemplate(w, "index.html", personData)
	if err != nil {
		http.Error(w, "internal server error", 500)
	}
}

func parseQueryParams(w http.ResponseWriter, r * http.Request) {
	v := r.FormValue("q")

	s := fmt.Sprintf("Your query param q: %v\nhowever if there is a " +
		"q in the body of a form it will overwrite the query param q", v)

	w.Header().Set("Content-Type", "text/text; charset=utf-8")
	_, err := w.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}
}
