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
	http.HandleFunc("/", parseForm)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseForm(w http.ResponseWriter, r *http.Request) {
	// read data from the body
	bs := make([]byte, r.ContentLength)  // length and cap = content length
	var b string
	var s string

	if r.Method == "POST" {
		_, err := r.Body.Read(bs)
		if err != nil {
			fmt.Print("BODY EOF")
		}
		s = string(bs)
		b = fmt.Sprintf("%b", bs)
		fmt.Printf("\n\n[]binary:\n%b\n", bs)
		fmt.Printf("\n\n[]byte:\n%v\n", bs)
		fmt.Printf("\n\nstring:\n%s\n\n", s)
	}

	data := struct{
		Binary string
		Bytes []byte
		String string
	}{
		Binary: b,
		Bytes: bs,
		String: s,
	}

	err := tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}