package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
	router.Post("/welcome/", PostWelcomeHandler)

	// the router implements http.Handler
	// so you just need to feed it into
	// http.ListenAndServe, or any other handler
	// that takes handlers such as common golang
	// middlewares
	log.Fatal(http.ListenAndServe(":1234", router))
}

// PostWelcomeHandler - Is an Implementation of http.HandlerFunc
func PostWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	s := string(body[:])
	sanitizer := strings.NewReplacer("<payload>", "", "</payload>", "", "&quot;", "\"")
	fmt.Println(sanitizer.Replace(s))
	w.Write([]byte(sanitizer.Replace(s)))
}

// GetWelcomeHandler - Is an Implementation of http.HandlerFunc
func GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("welcome!"))
}
