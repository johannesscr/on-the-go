package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		including an index.html file is a special case where the only file
		that is visible is the index.html file. no other files will be made
		available.
	*/
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}