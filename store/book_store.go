package store

import (
	"books-api/model"
	"books-api/util"
	"sync"
)

type BookStore struct {
	books map[string]model.Book
	mu    sync.RWMutex
}

func NewBookStore() *BookStore {
	return &BookStore{books: make(map[string]model.Book)}
}

func (store *BookStore) CreateBook(book model.Book) (model.Book, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, exists := store.books[book.ISBN]; exists {
		return model.Book{}, util.ErrBooksExists
	}

	store.books[book.ISBN] = book
	return book, nil
}

func (store *BookStore) UpdateBook(isbn string, bookDto model.Book) (model.Book, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	book, exists := store.books[isbn]
	if !exists {
		return model.Book{}, util.ErrBooksNotFound
	}
	book.ISBN = bookDto.ISBN
	book.Title = bookDto.Title
	book.Author = bookDto.Author
	book.ReleaseDate = bookDto.ReleaseDate

	store.books[isbn] = book
	return book, nil
}

func (store *BookStore) DeleteBook(isbn string) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	var book model.Book

	book, exists := store.books[isbn]
	if !book.Status || !exists {
		return util.ErrBooksNotFound
	}

	book.Status = false
	store.books[isbn] = book
	return nil
}

func (store *BookStore) GetAllBook() []model.Book {
	store.mu.Lock()
	defer store.mu.Unlock()

	books := make([]model.Book, 0, len(store.books))

	for _, book := range store.books {
		if book.Status {
			books = append(books, book)
		}
	}

	return books
}

func (store *BookStore) GetBookByISBN(isbn string) (model.Book, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	book, found := store.books[isbn]

	if !book.Status || !found {
		return model.Book{}, util.ErrBooksNotFound
	}

	return book, nil
}

func (store *BookStore) RestoreBookByISBN(isbn string) (model.Book, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	book, found := store.books[isbn]

	if !found {
		return model.Book{}, util.ErrBooksNotFound
	}

	if book.Status {
		return model.Book{}, util.ErrBookIsActive
	}

	book.Status = true
	store.books[isbn] = book

	return book, nil
}
