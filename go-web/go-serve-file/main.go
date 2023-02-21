package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/toby.png", dogIOCopy)
	http.HandleFunc("/toby2.png", dogServeContent)
	http.HandleFunc("/toby3.png", dogServeFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `
	<img width="300px" src="/toby.png" />
	<img width="300px" src="/toby2.png" />
	<img width="300px" src="/toby3.png" />
	`)
	if err != nil {
		log.Fatalln(err)
	}
}

func dogIOCopy(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer func(f *os.File){
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	_, err = io.Copy(w, f)
	if err != nil {
		log.Fatalln(err)
	}
}

func dogServeContent(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby2.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer func(f *os.File){
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}

func dogServeFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby3.png")
}
