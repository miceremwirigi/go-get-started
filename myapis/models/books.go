package models

type Book struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Publisher string  `json:"publisher"`
	Price     float64 `json:"price"`
	Pages     int     `json:"pages"`
}