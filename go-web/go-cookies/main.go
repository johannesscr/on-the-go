package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*html"))
}

func main() {
	http.HandleFunc("/", readCookie)
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/delete", deleteCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookie-name")
	if err == http.ErrNoCookie {
		c1 = &http.Cookie{
			Name: "my-cookie-name",
			Value: "0",
			MaxAge: 0, // MaxAge=0 means no 'Max-Age' attribute specified
		}
		http.SetCookie(w, c1)
	}

	c2, err := r.Cookie("visits")
	if err == http.ErrNoCookie {
		c2 = &http.Cookie{
			Name: "visits",
			Value: "0",
			MaxAge: 0, // MaxAge=0 means no 'Max-Age' attribute specified
		}
		http.SetCookie(w, c2)
	}

	fmt.Printf("c1 %v\n", c1)
	fmt.Printf("c2 %v\n", c2)

	i, _ := strconv.Atoi(c2.Value)
	i++ // increment visit count
	c2.Value = strconv.Itoa(i)
	http.SetCookie(w, c2)

	data := struct{
		CookieName *http.Cookie
		CookieVisit *http.Cookie
	}{
		CookieName: c1,
		CookieVisit: c2,
	}

	err = tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie-name",
		Value: "some random value",
		MaxAge: 60 * 60 * 24, // number of seconds the cookie lasts => 1 day
	})

	http.Redirect(w, r, "/", 307)
}

func deleteCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie-name")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", 307)
	}

	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", 307)
}
