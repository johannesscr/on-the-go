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
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/bar2", bar2)
	http.HandleFunc("/bar3", bar3)
	http.HandleFunc("/bar4", bar4)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at foo: %v\n\n", r.Method)
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at bar: %v\n\n", r.Method)
	w.Header().Set("Location", "/barred")
	w.WriteHeader(http.StatusSeeOther)
}

func bar2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at bar2: %v\n\n", r.Method)
	// or
	http.Redirect(w, r, "/barred", http.StatusSeeOther)
}

func bar3(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at bar3: %v\n\n", r.Method)
	w.Header().Set("Location", "/barred")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func bar4(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at bar3: %v\n\n", r.Method)
	w.Header().Set("Location", "/barred")
	w.WriteHeader(http.StatusMovedPermanently)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at barred: %v\n\n", r.Method)
}