package main

import (
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	log.Fatal(http.Serve(autocert.NewListener("example.com"), nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there: Your config: %+v", r.TLS)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}