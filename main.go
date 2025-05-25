package main

import (
	"books-api/handler"
	"books-api/logger"
	"books-api/model"
	"books-api/service"
	"books-api/store"
	"log"
	"net/http"
	"strings"
)

func main() {
	logger := logger.NewLogger()
	store := store.NewBookStore()
	seedBooks(store)
	service := service.NewBookService(store, logger)
	handler := handler.NewBookHandler(service)

	http.HandleFunc("/books", func(write http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			handler.CreateBookHandler(write, request)
		case http.MethodGet:
			handler.GetAllBooksHandler(write, request)
		default:
			http.Error(write, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/books/", func(write http.ResponseWriter, request *http.Request) {
		path := request.URL.Path

		if strings.HasPrefix(path, "/books/restore/") {
			if request.Method == http.MethodPut {
				handler.RestoreBookHandler(write, request)
			} else {
				http.Error(write, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
			return
		}

		switch request.Method {
		case http.MethodGet:
			handler.GetBookByISBNHandler(write, request)
		case http.MethodPut:
			handler.UpdateBookHandler(write, request)
		case http.MethodDelete:
			handler.DeleteBookHandler(write, request)
		default:
			http.Error(write, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	go logger.Start()
	logger.Log("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func seedBooks(store *store.BookStore) {
	books := []model.Book{
		{Title: "The Psychology of Money", Author: "Morgan Housel", ISBN: "999111222333", ReleaseDate: "2012-11-22", Status: true},
		{Title: "The Art Of Stoicism", Author: "Adora Kinara", ISBN: "999222333444", ReleaseDate: "2023-01-01", Status: true},
		{Title: "The Simplicity Pricinple", Author: "Julia Hobsbawm", ISBN: "999222333555", ReleaseDate: "2020-02-01", Status: true},
		{Title: "The Intelligent Investor", Author: "Benjamin Graham", ISBN: "999222333666", ReleaseDate: "1949-10-31", Status: true},
		{Title: "Rich Dad Poor Dad", Author: "Robert T. Kiyosaki & Sharon Lechter", ISBN: "999222333777", ReleaseDate: "1997-07-08", Status: true},
	}
	for _, book := range books {
		_, _ = store.CreateBook(book)
	}
}
