package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct (Model)

func main() {
	// init router
	r := mux.NewRouter()

	// route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// run server
	// takes in port and router
	log.Fatal(http.ListenAndServe(":8000", r))

}
