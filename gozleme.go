package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"

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
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	s := string(body[:])
	sanitizer := strings.NewReplacer("<payload>", "", "</payload>", "", "&quot;", "\"")
	sanitized := sanitizer.Replace(s)
	fmt.Println(sanitized)

	var m map[string]interface{}
	err = json.Unmarshal([]byte(sanitized), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("jenkins")
	err = c.Insert(m)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(sanitized))
}

// GetWelcomeHandler - Is an Implementation of http.HandlerFunc
func GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("welcome!"))
}
