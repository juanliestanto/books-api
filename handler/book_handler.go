package handler

import (
	"books-api/model"
	"books-api/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	services *service.BookService
}

func NewBookHandler(services *service.BookService) *BookHandler {
	return &BookHandler{services}
}

func responseJSON(write http.ResponseWriter, code int, message string, data interface{}) {
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(code)
	json.NewEncoder(write).Encode(map[string]interface{}{
		"status":  code,
		"message": message,
		"data":    data,
	})
}

func (handler *BookHandler) CreateBookHandler(write http.ResponseWriter, request *http.Request) {
	var book model.Book
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		http.Error(write, "Invalid JSON", http.StatusBadRequest)
		return
	}

	created, err := handler.services.Create(book)
	if err != nil {
		responseJSON(write, http.StatusConflict, err.Error(), nil)
		return
	}
	responseJSON(write, http.StatusCreated, "Book successfully created", created)
}

func (handler *BookHandler) GetAllBooksHandler(write http.ResponseWriter, request *http.Request) {
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(request.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	books := handler.services.GetAll(page, limit)

	responseJSON(write, http.StatusOK, "Get All Book successfully", books)
}

func (handler *BookHandler) GetBookByISBNHandler(write http.ResponseWriter, request *http.Request) {
	isbn := strings.TrimPrefix(request.URL.Path, "/books/")
	book, err := handler.services.GetByISBN(isbn)
	if err != nil {
		responseJSON(write, http.StatusNotFound, err.Error(), nil)
		return
	}
	responseJSON(write, http.StatusOK, "Get Book successfully", book)
}

func (handler *BookHandler) UpdateBookHandler(write http.ResponseWriter, request *http.Request) {
	isbn := strings.TrimPrefix(request.URL.Path, "/books/")
	var book model.Book
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		http.Error(write, "Invalid JSON", http.StatusBadRequest)
		return
	}
	updated, err := handler.services.UpdateByIsbn(isbn, book)
	if err != nil {
		responseJSON(write, http.StatusNotFound, err.Error(), nil)
		return
	}
	responseJSON(write, http.StatusOK, "Book successfully updated", updated)
}

func (handler *BookHandler) DeleteBookHandler(write http.ResponseWriter, request *http.Request) {
	isbn := strings.TrimPrefix(request.URL.Path, "/books/")

	err := handler.services.Delete(isbn)

	if err != nil {
		responseJSON(write, http.StatusNotFound, err.Error(), nil)
		return
	}
	responseJSON(write, http.StatusOK, "Book successfully deleted", nil)
}

func (handler *BookHandler) RestoreBookHandler(write http.ResponseWriter, request *http.Request) {
	isbn := strings.TrimPrefix(request.URL.Path, "/books/restore/")

	restore, err := handler.services.RestoreByIsbn(isbn)
	if err != nil {
		responseJSON(write, http.StatusNotFound, err.Error(), nil)
		return
	}
	responseJSON(write, http.StatusOK, "Book successfully restore", restore)
}
