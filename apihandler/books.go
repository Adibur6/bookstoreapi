package apihandler

import (
	"encoding/json"
	"github.com/adibur6/bookstoreapi/datahandler"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

// GetBooks returns all books in the BookList
func GetBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(datahandler.BookList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// BookGeneralized returns a list of all book names separated by newlines
func BookGeneralized(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	var bookNames []string
	for _, book := range datahandler.BookList {
		bookNames = append(bookNames, book.Name)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strings.Join(bookNames, "\n")))
}

// GetSingleBook returns a single book by ISBN
func GetSingleBook(w http.ResponseWriter, r *http.Request) {
	ISBN := chi.URLParam(r, "ISBN")
	book, exists := datahandler.BookList[ISBN]

	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// NewBook adds a new book to the BookList and updates the AuthorList accordingly
func NewBook(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming book data
	var newBook datahandler.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Validate the book data
	if newBook.ISBN == "" {
		http.Error(w, "ISBN cannot be empty", http.StatusBadRequest)
		return
	}

	if newBook.Name == "" {
		http.Error(w, "Book name cannot be empty", http.StatusBadRequest)
		return
	}

	if len(newBook.Authors) == 0 {
		http.Error(w, "There should be at least one author", http.StatusBadRequest)
		return
	}

	for _, author := range newBook.Authors {
		if author.Name == "" {
			http.Error(w, "Author name cannot be empty", http.StatusBadRequest)
			return
		}
	}

	// Add the new book to BookList
	datahandler.BookList[newBook.ISBN] = newBook

	// Update the AuthorList
	for _, author := range newBook.Authors {
		normalizedAuthorName := datahandler.SmStr(author.Name)
		authorBooks, exists := datahandler.AuthorList[normalizedAuthorName]

		if exists {
			// If the author already exists, just add the new book's ISBN to their list
			authorBooks.Books = append(authorBooks.Books, newBook.ISBN)
		} else {
			// If the author does not exist, create a new entry in AuthorList
			authorBooks = datahandler.AuthorBooks{
				Author: author,
				Books:  []string{newBook.ISBN},
			}
		}

		// Update the AuthorList with the new or modified author data
		datahandler.AuthorList[normalizedAuthorName] = authorBooks
	}

	// Respond with the status of the operation
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New book added successfully"))
}

// DeleteBook removes a book by ISBN and updates the AuthorList accordingly
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Get the ISBN from the URL parameters
	ISBN := chi.URLParam(r, "ISBN")

	// Check if the book exists in BookList
	book, exists := datahandler.BookList[ISBN]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Remove the book from BookList
	delete(datahandler.BookList, ISBN)

	// Remove the book's ISBN from the AuthorList
	for _, author := range book.Authors {
		normalizedAuthorName := datahandler.SmStr(author.Name)
		authorBooks, exists := datahandler.AuthorList[normalizedAuthorName]
		if exists {
			// Remove the book's ISBN from the author's list of books
			updatedBooks := []string{}
			for _, bISBN := range authorBooks.Books {
				if bISBN != ISBN {
					updatedBooks = append(updatedBooks, bISBN)
				}
			}
			// Update the AuthorList if there are remaining books
			if len(updatedBooks) > 0 {
				authorBooks.Books = updatedBooks
				datahandler.AuthorList[normalizedAuthorName] = authorBooks
			} else {
				// If no books remain for the author, remove the author from AuthorList
				delete(datahandler.AuthorList, normalizedAuthorName)
			}
		}
	}

	// Respond with the status of the operation
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully"))
}

// UpdateBook updates an existing book by ISBN and modifies its authors accordingly
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Get the ISBN from the URL parameters
	ISBN := chi.URLParam(r, "ISBN")

	// Decode the incoming book data
	var updatedBook datahandler.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Validate the updated book data
	if updatedBook.ISBN == "" {
		http.Error(w, "ISBN cannot be empty", http.StatusBadRequest)
		return
	}

	if updatedBook.Name == "" {
		http.Error(w, "Book name cannot be empty", http.StatusBadRequest)
		return
	}

	if len(updatedBook.Authors) == 0 {
		http.Error(w, "There should be at least one author", http.StatusBadRequest)
		return
	}

	for _, author := range updatedBook.Authors {
		if author.Name == "" {
			http.Error(w, "Author name cannot be empty", http.StatusBadRequest)
			return
		}
	}

	// Check if the book exists in BookList
	oldBook, exists := datahandler.BookList[ISBN]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Remove the old book from AuthorList
	for _, author := range oldBook.Authors {
		normalizedAuthorName := datahandler.SmStr(author.Name)
		authorBooks, exists := datahandler.AuthorList[normalizedAuthorName]
		if exists {
			// Remove the book's ISBN from the author's list of books
			updatedBooks := []string{}
			for _, bISBN := range authorBooks.Books {
				if bISBN != ISBN {
					updatedBooks = append(updatedBooks, bISBN)
				}
			}
			if len(updatedBooks) > 0 {
				authorBooks.Books = updatedBooks
				datahandler.AuthorList[normalizedAuthorName] = authorBooks
			} else {
				delete(datahandler.AuthorList, normalizedAuthorName)
			}
		}
	}

	// Update the BookList with new book details
	datahandler.BookList[ISBN] = updatedBook

	// Add the new authors to AuthorList
	for _, author := range updatedBook.Authors {
		normalizedAuthorName := datahandler.SmStr(author.Name)
		authorBooks, exists := datahandler.AuthorList[normalizedAuthorName]

		if exists {
			// If the author already exists, add the new book's ISBN to their list
			authorBooks.Books = append(authorBooks.Books, ISBN)
		} else {
			// If the author does not exist, create a new entry in AuthorList
			authorBooks = datahandler.AuthorBooks{
				Author: author,
				Books:  []string{ISBN},
			}
		}
		datahandler.AuthorList[normalizedAuthorName] = authorBooks
	}

	// Respond with the status of the operation
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book updated successfully"))
}
