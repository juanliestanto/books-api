package util

import "errors"

var (
	ErrBooksExists   = errors.New("Book Already Exists")
	ErrBooksNotFound = errors.New("Book Not Found")
	ErrBookIsActive  = errors.New("Book Is Active")
)
