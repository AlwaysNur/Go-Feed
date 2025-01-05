package main

import (
	api "feed/api"
	staticserver "feed/static_server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Starting server on http://localhost:8080")
	fmt.Println("Starting API on http://localhost:8080/api")
	// main
	r.HandleFunc("/submit", staticserver.SubmitPage)
	r.HandleFunc("/css", staticserver.Css)
	r.HandleFunc("/", api.Response).Methods("GET")
	// api post, get
	r.HandleFunc("/api", api.HandlePost).Methods("POST")
	// host on port 8080
	log.Fatal(http.ListenAndServe(":8080", r))

}
