package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Items []string `json:"items"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/unmarshal", unmarshal)
	http.HandleFunc("/json", jsonPage)
	http.HandleFunc("/json-body", jsonBody)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Please use routes /enc /marsh"))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func marshal(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		FirstName: "james",
		LastName: "bond",
		Items: []string{"suit", "gun", "glasses"},
	}

	// json marshal
	jsonData, err := json.Marshal(p1)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func encode(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		FirstName: "james",
		LastName: "bond",
		Items: []string{"suit", "gun", "glasses"},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func unmarshal(w http.ResponseWriter, r *http.Request) {
	jsonData := `{
    "first_name": "miss",
    "last_name": "moneypenny",
    "items": ["lipstick", "dress", "glasses", "gloves"]
}`
	var p1 Person

	err := json.Unmarshal([]byte(jsonData), &p1)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Println(p1)

	_, err = w.Write([]byte("unmarshal done"))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func jsonPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func readBody(w http.ResponseWriter, r *http.Request) []byte {
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 400)
	}
	log.Printf("body: %s", bs)
	return bs
}

func parseBody(w http.ResponseWriter, r *http.Request, i interface{}) {
	bs := readBody(w, r)
	log.Print("parseBody interface ", i)
	err := json.Unmarshal(bs, i)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 400)
	}
}

func jsonBody(w http.ResponseWriter, r *http.Request) {
	// location to store data received
	var p1 Person
	parseBody(w, r, &p1)
	log.Println(p1)

	// send json back
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}