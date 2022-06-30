package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author Author `json:"author"`
}

type Author struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

var books []Book

func dummyDataLoad()  {
	books = append(books,
		Book{
			ID:   "1",
			Name: "Book 1",
			Author: Author{
				Name: "Abul Karim",
				ID:   "1",
			},
		},
	)

	books = append(books,
		Book{
			ID:   "2",
			Name: "Book 2",
			Author: Author{
				Name: "Abul Bashar",
				ID:   "2",
			},
		},
	)

	books = append(books,
		Book{
			ID:   "3",
			Name: "Book 3",
			Author: Author{
				Name: "Abul Kashem",
				ID:   "3",
			},
		},
	)
}

func main() {

	dummyDataLoad()

	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/show/{id}", showBook).Methods("GET")

	http.ListenAndServe(":2001", r)
}

func getBooks(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func showBook(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, book := range books {
		if book.ID == params["id"]{
			json.NewEncoder(writer).Encode(book)
			return
		}
	}

	json.NewEncoder(writer).Encode(Book{})
}
