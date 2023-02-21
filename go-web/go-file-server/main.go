package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/home", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `
	<img width="400px" src="/assets/img/toby.png" />
	`)
	if err != nil {
		http.Error(w, "dir not found", 404)
	}
}
