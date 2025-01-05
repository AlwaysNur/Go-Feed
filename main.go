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
	r.HandleFunc("/", staticserver.Index)
	r.HandleFunc("/submit", staticserver.SubmitPage)
	r.HandleFunc("/css", staticserver.Css)
	// api post, get
	r.HandleFunc("/api", api.HandlePost).Methods("POST")
	r.HandleFunc("/api", api.Response).Methods("GET")
	// host on port 8080
	log.Fatal(http.ListenAndServe(":8080", r))

}
