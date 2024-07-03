package search

import (
	"fmt"
	"testing"
)

// //////////////////////////////////////////////////// //
// Tests are NOT mocked, just for development purposes  //
// //////////////////////////////////////////////////// //
func Test(t *testing.T) {
	// setup
	search := NewSearch(&SearchConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "masterKey",
	})
	// index create
	search.GetOrCreateIndex("movies")
	// uplaod documents
	documents := []map[string]interface{}{
		{"id": 1, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 2, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 3, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 4, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 5, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 6, "title": "Philadelphia", "genres": []string{"Drama"}},
	}
	search.UploadDocuments("movies", documents)
	// search
	res, _ := search.Search("movies", "carol", 10)
	fmt.Println(res)

	// delete document
	search.DeleteDocument("movies", "2")

	// delete all documents
	search.DeleteAllDocuments("movies")

	// delete index
	search.DeleteIndex("movies")
}
