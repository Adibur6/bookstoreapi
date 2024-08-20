package apihandler

import (
	"encoding/json"
	"github.com/adibur6/bookstoreapi/datahandler"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// GetAuthors returns all authors and their associated books
func GetAuthors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(datahandler.AuthorList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetSingleAuthor returns a single author by name along with their associated books
func GetSingleAuthor(w http.ResponseWriter, r *http.Request) {
	authorName := chi.URLParam(r, "AuthorName")
	normalizedAuthorName := datahandler.SmStr(authorName)
	authorBooks, exists := datahandler.AuthorList[normalizedAuthorName]

	if !exists {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(authorBooks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
