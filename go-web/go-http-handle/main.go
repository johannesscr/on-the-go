package main

import (
	"log"
	"net/http"
)

type hotdog int
func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Any code you want\n"))
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	var d hotdog
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Panic(err)
	}
}