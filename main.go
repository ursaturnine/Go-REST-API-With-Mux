package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct (Model)
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init Books variable as a slice Book struct
var books []Book

// Get all Books
// route handlers must take in response and request
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init router
	r := mux.NewRouter()

	// Mock Data
	// to do - implement database
	books = append(books, Book{ID: "1", Isbn: "448734", Title: "Book1", Author: &Author{
		FirstName: "Olivia", LastName: "Kaye",
	}})
	books = append(books, Book{ID: "2", Isbn: "448735", Title: "Book2", Author: &Author{
		FirstName: "Tyrah", LastName: "Denise",
	}})
	books = append(books, Book{ID: "3", Isbn: "448736", Title: "Book3", Author: &Author{
		FirstName: "Aliyah", LastName: "Renee",
	}})
	books = append(books, Book{ID: "4", Isbn: "448737", Title: "Book4", Author: &Author{
		FirstName: "Vanessa", LastName: "Ragine",
	}})
	books = append(books, Book{ID: "5", Isbn: "448738", Title: "Book5", Author: &Author{
		FirstName: "Natalie", LastName: "Monet",
	}})

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
