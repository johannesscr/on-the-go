package main

import (
	"encoding/json"
	"github.com/johannesscr/go-web-basics/go-postgres/models"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/*
		curl http://localhost:8080/books
		curl http://localhost:8080/books?isbn=978-1505255607
	*/
	http.HandleFunc("/books", BooksIndex)
	/*
		curl -X POST http://localhost:8080/book -H "Content-Type:application/json" -d '{"isbn": "978-1501234567", "title": "new", "author": "mr auth", "price": 2.13}'
	*/
	http.HandleFunc("/book", BooksCreate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func BooksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		code := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(code), code)
	}

	isbn := r.URL.Query().Get("isbn")
	log.Println("isbn:", isbn)
	if isbn != "" {
		if book, ok := models.SelectBook(isbn); ok {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				code := http.StatusInternalServerError
				http.Error(w, http.StatusText(code), code)
			}
			return
		}
	}

	if books, ok := models.SelectBooks(); ok {
		err := json.NewEncoder(w).Encode(books)
		if err != nil {
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
		}
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

func ParseBody(w http.ResponseWriter, r *http.Request, i interface{}) {
	bs := readBody(w, r)
	log.Print("parseBody interface ", i)
	err := json.Unmarshal(bs, i)
	log.Print("unmarshalled interface ", i)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 400)
	}
}

func BooksCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		code := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(code), code)
	}

	bk := models.Book{}
	ParseBody(w, r, &bk)
	if ok := models.InsertBook(&bk); ok {
		http.Redirect(w, r, "/books", http.StatusSeeOther)
		return
	}

	http.Error(w, "internal server error", 500)
}
