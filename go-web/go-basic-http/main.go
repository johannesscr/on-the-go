package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type hotdog int
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct{
		Method string
		URL *url.URL
		Form map[string][]string
		Header http.Header
		Host string
		ContentLength int64
	}{
		Method: r.Method,
		URL: r.URL,
		Form: r.Form,
		Header: r.Header,
		Host: r.Host,
		ContentLength: r.ContentLength,
	}

	err = tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}

type boerierol int
func (m boerierol) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Token", "my special token")

	s := "any code here\n"
	_, err := w.Write([]byte(s))
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
}

func main() {
	var d hotdog
	var g boerierol

	s := "response"
	switch s {
	case "request":
		log.Fatalln(http.ListenAndServe(":8080", d))
	case "response":
		log.Fatalln(http.ListenAndServe(":8080", g))
	}
}