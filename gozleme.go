package main

import (
	"log"
	"net/http"

	"gopkg.in/husobee/vestigo.v1"
)

func main() {
	// create a new vestigo router
	router := vestigo.NewRouter()

	// Define a GET /welcome route, and
	// specify the standard http.HandlerFunc
	// to run from that route
	router.Get("/welcome", GetWelcomeHandler)

	// Define a POST /welcome/:name where :name
	// is a URL parameter that will be available
	// anywhere you have access to the http.Request
	router.Post("/welcome/:name", PostWelcomeHandler)

	// the router implements http.Handler
	// so you just need to feed it into
	// http.ListenAndServe, or any other handler
	// that takes handlers such as common golang
	// middlewares
	log.Fatal(http.ListenAndServe(":1234", router))
}

// PostWelcomeHandler - Is an Implementation of http.HandlerFunc
func PostWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	name := vestigo.Param(r, "name") // url params live in the request
	w.WriteHeader(200)
	w.Write([]byte("welcome " + name + "!"))
}

// GetWelcomeHandler - Is an Implementation of http.HandlerFunc
func GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("welcome!"))
}
