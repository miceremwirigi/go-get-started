package main

import (
	"fmt"
	"go-getting-started/myapis/handlers"
	"go-getting-started/myapis/models"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Hello")
}

func main() {

	handlers.Books = []models.Book{
		{ID: uuid.New().String(), Title: "Merlin", Publisher: "Arthur", Price: 100.0, Pages: 1000},
		{ID: uuid.New().String(), Title: "Blacklist", Publisher: "Red", Price: 300.0, Pages: 640},
		{ID: uuid.New().String(), Title: "Shadow", Publisher: "Hunter", Price: 150.0, Pages: 800},
		{ID: uuid.New().String(), Title: "Avatar", Publisher: "Aang", Price: 90.0, Pages: 1300},
	}

	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/books/", handlers.BookHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
