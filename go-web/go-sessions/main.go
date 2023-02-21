package main

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	UserName       string
	First          string
	Last           string
	hashedPassword string
}

var tpl *template.Template
var dbSession = make(map[string]string, 1)
var dbUser = make(map[string]user, 1)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login-page", loginPage)
	http.HandleFunc("/info-page", infoPage)
	http.HandleFunc("/register-page", registerPage)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := getSession(w, r)
	username, ok := dbSession[cookie.Value]
	if ok {
		u := dbUser[username]
		err := tpl.ExecuteTemplate(w, "index.html", u)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func infoPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := getSession(w, r)
	username, ok := dbSession[cookie.Value]
	if ok {
		u := dbUser[username]
		err := tpl.ExecuteTemplate(w, "info.html", u)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		return
	}
	http.Redirect(w, r, "/", 307)
}

func getSession(w http.ResponseWriter, r *http.Request) (*http.Cookie, bool) {
	cookie, err := r.Cookie("session")
	if err != nil {
		//http.Error(w, err.Error(), 500)
		return &http.Cookie{}, false
	}
	return cookie, true
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	// check if user exists
	u, ok := dbUser[username]
	fmt.Printf("username: %v\nuser: %v\nok: %v\n",
		username, u, ok)
	if !ok {
		http.Redirect(w, r, "/register-page", 307)
		return
	}

	// check password
	hashedPasswordBs := []byte(u.hashedPassword)
	passwordBs := []byte(r.FormValue("password"))
	valid := bcrypt.CompareHashAndPassword(hashedPasswordBs, passwordBs)
	if valid != nil {
		http.Redirect(w, r, "/login-page", 307)
		return
	}

	// set cookie
	UUID := uuid.New()  // in go keep acronyms uppercase
	fmt.Printf("uuid: %T %v", UUID, UUID)
	cookie := &http.Cookie{
		Name:  "session",
		Value: UUID.String(),
		//Secure: true, // for HTTPS only
		HttpOnly: true, // only accessible by HTTP not JavaScript
	}

	// set session
	dbSession[cookie.Value] = u.UserName

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 307)
}

func logout(w http.ResponseWriter, r *http.Request) {
	// get cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// delete
	delete(dbSession, cookie.Value)

	// destroy cookie
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 307)
}

func register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	u, ok := dbUser[username]
	if !ok {
		hashBs, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 4)
		hash := string(hashBs)
		u = user{
			UserName: r.FormValue("username"),
			First:    r.FormValue("first"),
			Last:     r.FormValue("last"),
			hashedPassword: hash,
		}
		// add user to db
		dbUser[u.UserName] = u
	}
	http.Redirect(w, r, "/", 307)
}
