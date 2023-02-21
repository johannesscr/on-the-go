package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", handler()))
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/double", doubleHandler)
	return r
}

func doubleHandler(w http.ResponseWriter, r*http.Request) {
	text := r.URL.Query().Get("v")
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}
	v, err := strconv.Atoi(text) // convert string to int
	if err != nil {
		http.Error(w, "not a number", http.StatusBadRequest)
		return
	}

	_, err = fmt.Fprintf(w, "%v", v*2)
	if err != nil {
		log.Println(err)
	}
}

