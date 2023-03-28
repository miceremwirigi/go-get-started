package handlers

import (
	"encoding/json"
	"go-getting-started/myapis/models"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var Books = []models.Book{}

func BookHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		id := strings.TrimPrefix(r.URL.Path, "/books/")

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
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var book models.Book
		err = json.Unmarshal(body, &book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		book.ID = uuid.New().String()
		Books = append(Books, book)
		json.NewEncoder(w).Encode(Books)
		return

	} else if r.Method == "PUT" {
		id := strings.TrimPrefix(r.URL.Path, "/books/")

		if id == "" {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		for count, book := range Books {
			if book.ID == id {
				// log.Println("Found book with id: " + book.ID)
				var result models.Book
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
				Books[count] = book

				json.NewEncoder(w).Encode(Books)
				return
			}
		}
		http.Error(w, "invalid id could not find book", http.StatusBadRequest)
		return

	} else if r.Method == "DELETE" {
		id := strings.TrimPrefix(r.URL.Path, "/books/")
		var newBooks []models.Book
		valid := false

		if id == "" {
			http.Error(w, "empyty id", http.StatusBadRequest)
			return
		}

		for _, book := range Books {
			if book.ID == id {
				valid = true
			} else {
				newBooks = append(newBooks, book)
			}
		}

		if !valid {
			http.Error(w, "invalid id could not find book", http.StatusBadRequest)
			return
		}

		Books = newBooks
		encoder := json.NewEncoder(w)
		encoder.Encode(Books)
		return
	}
}
