package util

import "errors"

var (
	ErrBooksExists   = errors.New("Book Already Exists")
	ErrBooksNotFound = errors.New("Book Not Found")
	ErrBookIsActive  = errors.New("Book Is Active")
	ErrIsbnNotMatch  = errors.New("ISBN In URL Parameter Does Not Match ISBN In Request Body")
)
