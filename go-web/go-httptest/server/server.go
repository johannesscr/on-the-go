package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// upperCaseHandler get the query param "word" and transforms it to uppercase
// Example:
// 	request: http://localhost:1234/upper?word=abc
// 	response: ABC
func upperCaseHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "invalid request")
		return
	}
	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "missing word")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, strings.ToUpper(word))
}

func main() {
	http.HandleFunc("/upper", upperCaseHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
