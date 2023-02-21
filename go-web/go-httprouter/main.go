package main

import (
	"github.com/johannesscr/go-web-basics/go-httprouter/routes"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", routes.Home)
	router.GET("/hello/:name", routes.Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
