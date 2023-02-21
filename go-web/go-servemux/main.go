package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("home"))
	if err != nil {
		log.Fatalln(err)
	}
}

//func main() {
//  // using the Go serve mux
//  // 3rd party should be similar
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", home)
//	log.Fatal(http.ListenAndServe(":8080", mux))
//}

func main() {

	http.HandleFunc("/", home)
	// using the default serve mux
	log.Fatal(http.ListenAndServe(":8080", nil))
}