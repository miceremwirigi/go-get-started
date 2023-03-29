package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-getting-started/myapis/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TearDown() {
	Books = []models.Book{}
}

func TestGetBookHandler(t *testing.T) {
	url := "/books/"

	Books = []models.Book{
		{ID: "01", Title: "Book1", Publisher: "Publisher1", Price: 200, Pages: 120},
		{ID: "02", Title: "Book2", Publisher: "Publisher2", Price: 300, Pages: 140},
		{ID: "03", Title: "Book3", Publisher: "Publisher3", Price: 400, Pages: 160},
		{ID: "04", Title: "Book4", Publisher: "Publisher4", Price: 500, Pages: 180},
	}

	defer TearDown()

	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)
	req.Header.Set("content-type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BookHandler)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var result []models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, len(Books), len(result))
}

func TestGetSpecificBookHandler(t *testing.T) {
	url := "/books/02"

	Books = []models.Book{
		{ID: "01", Title: "Book1", Publisher: "Publisher1", Price: 200, Pages: 120},
		{ID: "02", Title: "Book2", Publisher: "Publisher2", Price: 300, Pages: 140},
		{ID: "03", Title: "Book3", Publisher: "Publisher3", Price: 400, Pages: 160},
		{ID: "04", Title: "Book4", Publisher: "Publisher4", Price: 500, Pages: 180},
	}

	defer TearDown()

	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)
	req.Header.Set("content-type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BookHandler)

	handler.ServeHTTP(rr, req)

	var result models.Book
	body, err := io.ReadAll(rr.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(body, &result)
	assert.Nil(t, err)
	assert.Equal(t, result.Title, "Book2")

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPostBookHandler(t *testing.T) {
	api := "/books/"
	data := []byte(`{
		"title":     "title",
		"publisher": "publisher",
		"price":     200,
		"pages":     433
	}`)

	Books = []models.Book{}

	defer TearDown()

	// book := &models.Book{Title: "title", Publisher: "publisher", Price: 200, Pages: 210}
	// body, err := json.Marshal(book)
	// if err != nil {
	// 	t.Errorf("got error marshaling book: %v", err)
	// }
	// req, err := http.NewRequest("POST", api, bytes.NewBuffer(body))

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(data))
	assert.Nil(t, err)
	req.Header.Set("content-type", "application/json")

	handler := http.HandlerFunc(BookHandler)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, rr.Code)
	}
	assert.Equal(t, http.StatusOK, rr.Code)

	var result []models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	assert.Nil(t, err)
	fmt.Println(result)

	assert.Equal(t, len(Books), len(result))
	assert.True(t, len(result) == 1)
}

func TestPutBookHandler(t *testing.T) {
	api := "/books/01"
	data := []byte(`{
		"title": "English Aid",
		"publisher": "Moran",
		"price": 320,
		"pages": 213
	}`)

	Books = []models.Book{
		{ID: "01", Title: "Book1", Publisher: "Publisher1", Price: 200, Pages: 120},
		{ID: "02", Title: "Book2", Publisher: "Publisher2", Price: 300, Pages: 140},
		{ID: "03", Title: "Book3", Publisher: "Publisher3", Price: 400, Pages: 160},
		{ID: "04", Title: "Book4", Publisher: "Publisher4", Price: 500, Pages: 180},
	}

	defer TearDown()

	req, err := http.NewRequest("PUT", api, bytes.NewBuffer(data))
	assert.Nil(t, err)

	req.Header.Set("content-type", "application/json")

	handler := http.HandlerFunc(BookHandler)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result []models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, len(Books), len(result))
}

func TestDeleteBookHandler(t *testing.T) {
	api := "/books/01"
	Books = []models.Book{
		{ID: "01", Title: "Book1", Publisher: "Publisher1", Price: 200, Pages: 120},
		{ID: "02", Title: "Book2", Publisher: "Publisher2", Price: 300, Pages: 140},
		{ID: "03", Title: "Book3", Publisher: "Publisher3", Price: 400, Pages: 160},
		{ID: "04", Title: "Book4", Publisher: "Publisher4", Price: 500, Pages: 180},
	}

	defer TearDown()

	req, err := http.NewRequest("DELETE", api, nil)
	assert.Nil(t, err)

	req.Header.Set("content-type", "application/json")

	handler := http.HandlerFunc(BookHandler)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result = []models.Book{}
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	assert.Nil(t, err)

	assert.Equal(t, len(Books), len(result))
}

// ----------------------DRY----------------------//
func TestBooksApi(t *testing.T) {

	Books = []models.Book{
		{ID: "01", Title: "Book1", Publisher: "Publisher1", Price: 200, Pages: 120},
		{ID: "02", Title: "Book2", Publisher: "Publisher2", Price: 300, Pages: 140},
		{ID: "03", Title: "Book3", Publisher: "Publisher3", Price: 400, Pages: 160},
		{ID: "04", Title: "Book4", Publisher: "Publisher4", Price: 500, Pages: 180},
	}

	data := []byte(`{
		"title":     "title",
		"publisher": "publisher",
		"price":     200,
		"pages":     433
	}`)

	BooksApiTests := []struct {
		description  string
		route        string
		expectedCode int
		httpMethod   string
		payload      []byte
	}{
		{
			description:  "test get all books",
			route:        "/books/",
			expectedCode: http.StatusOK,
			httpMethod:   "GET",
			payload:      nil,
		},
		{
			description:  "test get specific book",
			route:        "/books/02",
			expectedCode: http.StatusOK,
			httpMethod:   "GET",
			payload:      nil,
		},
		{
			description:  "test post book",
			route:        "/books/",
			expectedCode: http.StatusOK,
			httpMethod:   "POST",
			payload:      data,
		},
		{
			description:  "test put specific book",
			route:        "/books/04",
			expectedCode: http.StatusOK,
			httpMethod:   "PUT",
			payload:      data,
		},
		{
			description:  "test delete specific book",
			route:        "/books/01",
			expectedCode: http.StatusOK,
			httpMethod:   "DELETE",
			payload:      nil,
		},
	}

	for _, test := range BooksApiTests {
		req, err := http.NewRequest(test.httpMethod, test.route, bytes.NewBuffer(test.payload))
		assert.Nil(t, err)
		req.Header.Set("content-type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(BookHandler)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.expectedCode, rr.Code)

		if test.httpMethod == "GET" && test.description == "test get specific book" {
			var result models.Book
			body, err := io.ReadAll(rr.Body)
			assert.Nil(t, err)
			err = json.Unmarshal(body, &result)
			assert.Nil(t, err)

		} else {
			var result []models.Book
			err = json.Unmarshal(rr.Body.Bytes(), &result)
			assert.Nil(t, err)
		}

	}
}
