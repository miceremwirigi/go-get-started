package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	// "strconv"

	"github.com/google/uuid"
)

type Book struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Publisher string  `json:"publisher"`
	Price     float64 `json:"price"`
	Pages     int     `json:"pages"`
}

var Books = []Book{}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Hello")
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// log.Println("The path is: " + r.URL.Path)
		id := strings.TrimPrefix(r.URL.Path, "/books/")
		// log.Println("Query is: ", id)

		if id != "" {
			for _, book := range Books {
				if book.ID == id {
					// log.Println("Found book with id: " + book.ID)
					json.NewEncoder(w).Encode(book)
					return
				}
			}

			w.Write([]byte("{}"))
			return
		}

		json.NewEncoder(w).Encode(Books)
		return

	} else if r.Method == "POST" {

		body, err := io.ReadAll(r.Body)
		// log.Println("the query is: ", body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var book Book
		err = json.Unmarshal(body, &book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		book.ID = uuid.New().String()
		Books = append(Books, book)
		json.NewEncoder(w).Encode(Books)
	} else if r.Method == "PUT" {
		id := strings.TrimPrefix(r.URL.Path, "/books/")

		if id == "" {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		for _, book := range Books {
			if book.ID == id {
				log.Println("Found book with id: " + book.ID)
				var result Book
				body, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				err = json.Unmarshal(body, &result)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				book.Title = result.Title
				book.Publisher = result.Publisher
				book.Price = result.Price
				book.Pages = result.Pages

				json.NewEncoder(w).Encode(book)
				return
			}
		}
		http.Error(w, "invalid id could not find book", http.StatusBadRequest)
		return
	}

}

func main() {

	Books = []Book{
		{ID: uuid.New().String(), Title: "Merlin", Publisher: "Arthur", Price: 100.0, Pages: 1000},
		{ID: uuid.New().String(), Title: "Blacklist", Publisher: "Red", Price: 300.0, Pages: 640},
		{ID: uuid.New().String(), Title: "Shadow", Publisher: "Hunter", Price: 150.0, Pages: 800},
		{ID: uuid.New().String(), Title: "Avatar", Publisher: "Aang", Price: 90.0, Pages: 1300},
	}

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/books/", BookHandler)
	// http.HandleFunc("/books/:id", BookHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
