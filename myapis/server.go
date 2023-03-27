package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	// "strconv"
)

type Book struct {
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
		json.NewEncoder(w).Encode(Books)
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
		
		Books = append(Books, book)
		json.NewEncoder(w).Encode(Books)
	}

}

func main() {

	Books = []Book{
		{Title: "Merlin", Publisher: "Arthur", Price: 100.0, Pages: 1000},
		{Title: "Blacklist", Publisher: "Red", Price: 300.0, Pages: 640},
		{Title: "Shadow", Publisher: "Hunter", Price: 150.0, Pages: 800},
		{Title: "Avatar", Publisher: "Aang", Price: 90.0, Pages: 1300},
	}

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/books", BookHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
