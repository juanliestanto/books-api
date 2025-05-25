package model

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	ReleaseDate string `json:"release_date"`
	Status      bool   `json:"status"` //TODO THIS VARIABLE FOR SOFT DELETE, 0 ACTIVE 1 NON ACTIVE
}
