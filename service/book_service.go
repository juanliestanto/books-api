package service

import (
	"books-api/logger"
	"books-api/model"
	"books-api/store"
	"books-api/util"
)

type BookService struct {
	store  *store.BookStore
	logger *logger.Logger
}

func NewBookService(store *store.BookStore, logger *logger.Logger) *BookService {
	return &BookService{store, logger}
}

func (service *BookService) Create(book model.Book) (model.Book, error) {
	createBook, err := service.store.CreateBook(book)

	if err == nil {
		service.logger.Log("Book Created : " + createBook.ISBN)
	}

	return createBook, err
}

func (service *BookService) Delete(isbn string) error {
	err := service.store.DeleteBook(isbn)

	if err == nil {
		service.logger.Log("Book Deleted : " + isbn)
	}

	return err
}

func (service *BookService) UpdateByIsbn(isbn string, book model.Book) (model.Book, error) {
	if isbn != book.ISBN {
		return model.Book{}, util.ErrIsbnNotMatch
	}

	updateBook, err := service.store.UpdateBook(isbn, book)

	if err == nil {
		service.logger.Log("Book Updated : " + isbn)
	}

	return updateBook, err
}

func (service *BookService) RestoreByIsbn(isbn string) (model.Book, error) {
	restoreBook, err := service.store.RestoreBookByISBN(isbn)

	if err == nil {
		service.logger.Log("Book Restore : " + isbn)
	}

	return restoreBook, err
}

func (service *BookService) GetByISBN(isbn string) (model.Book, error) {
	getByIsbn, err := service.store.GetBookByISBN(isbn)

	if err == nil {
		service.logger.Log("Book Restore : " + isbn)
	}

	return getByIsbn, err
}

func (service *BookService) GetAll(page, limit int) []model.Book {
	books := service.store.GetAllBook()
	start := (page - 1) * limit
	end := start + limit
	if start > len(books) {
		return []model.Book{}
	}
	if end > len(books) {
		end = len(books)
	}
	return books[start:end]
}
