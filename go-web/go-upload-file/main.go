package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

type FileContent struct {
	Bytes []byte
	String string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*html"))
}

func main() {
	http.HandleFunc("/save", saveFile)
	http.HandleFunc("/serve/", readFile)
	http.HandleFunc("/read/txt", readFileTxt)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.Error(w, "<em>page not found</em>", 404)
	}
}

func readFileTxt(w http.ResponseWriter, r *http.Request) {
	fileContent := FileContent{}
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		fFile, fHeader, err := r.FormFile("uploadFile")  // name from the form
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func(f multipart.File){
			err := f.Close()
			if err != nil {

			}
		}(fFile)

		fmt.Printf("\nfile:%v\nheader: %v\nerror: %v\ncontent: %v\n",
			fFile,  fHeader.Header, err)

		bs, err := ioutil.ReadAll(fFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s := string(bs)
		fmt.Printf("### Print Bytes ###\n%v\n", bs)
		fmt.Printf("### Print String ###\n%v\n", s)

		fileContent = FileContent{
			Bytes: bs,
			String: s,
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "index.html", fileContent)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.Error(w, "<em>page not found</em>", 404)
	}
}

func saveFile(w http.ResponseWriter, r *http.Request) {
	fileContent := FileContent{}
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		fFile, fHeader, err := r.FormFile("uploadFile")  // name from the form
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func(f multipart.File){
			err := f.Close()
			if err != nil {

			}
		}(fFile)

		fmt.Printf("\nfile:%v\nheader: %v\nerror: %v\ncontent: %v\n",
			fFile,  fHeader.Header, err)

		bs, err := ioutil.ReadAll(fFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s := string(bs)
		fmt.Printf("### Print Bytes ###\n%v\n", bs)
		fmt.Printf("### Print String ###\n%v\n", s)

		dst, err := os.Create(filepath.Join("./documents/", fHeader.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer func(f *os.File){
			err := f.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}(dst)

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fileContent = FileContent{
			Bytes: bs,
			String: s,
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "index.html", fileContent)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.Error(w, "<em>page not found</em>", 404)
	}
}

func readFile(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("filename")
	fmt.Printf("filename: %v\n", filename)
	if filename == "" {
		http.Error(w, "valid filename", 400)
	}
	http.ServeFile(w, r, filepath.Join("./documents/", filename))
}