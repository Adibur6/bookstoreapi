package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"strings"
)

// Author struct holds common information of an Author
type Author struct {
	Name string `json:"name,omitempty"`
	Home string `json:"home"`
	Age  string `json:"age"`
}

// AuthorBooks Use Composition to store Books
// ISBN which can be different for each
type AuthorBooks struct {
	Author `json:"author"`
	Books  []string `json:"books"`
}

// Book store Book information and the Authors who authored it
type Book struct {
	Name    string   `json:"book_name,omitempty"`
	Authors []Author `json:"authors"`
	ISBN    string   `json:"isbn,omitempty"`
	Genre   string   `json:"genre"`
	Pub     string   `json:"publisher"`
}

// Credentials Stores Login Credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// BookDB AuthorDB CredDB are databases
type BookDB map[string]Book
type AuthorDB map[string]AuthorBooks
type CredDB map[string]string

// BookList AuthorList CredList are DB Instances
var BookList BookDB
var AuthorList AuthorDB
var CredList CredDB

func SmStr(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", ""))
}
func IntializeDB() {
	author1 := Author{
		Name: "Ashfaqur Rahman",
		Home: "America",
		Age:  "25",
	}
	author2 := Author{
		Name: "Monir Khan",
		Home: "Jamaica",
		Age:  "29",
	}
	author3 := Author{
		Name: "Jamal Kabir",
		Home: "India",
		Age:  "37",
	}
	book1 := Book{
		Name:    "The burning sun",
		Authors: []Author{author1, author2},
		ISBN:    "ISBN1",
		Genre:   "History",
		Pub:     "Newyork Times",
	}
	book2 := Book{
		Name:    "Smiling Fox",
		Authors: []Author{author1},
		ISBN:    "ISBN2",
		Genre:   "Satire",
		Pub:     "Eden Club",
	}
	book3 := Book{
		Name:    "Hunting The Crow",
		Authors: []Author{author3},
		ISBN:    "ISBN3",
		Genre:   "Adventure",
		Pub:     "Tigers Publication",
	}
	User := Credentials{
		Username: "adib",
		Password: "1234",
	}
	BookList = make(BookDB)
	AuthorList = make(AuthorDB)
	CredList = make(CredDB)

	var ab1 AuthorBooks
	ab1.Author = author1
	ab1.Books = append(ab1.Books, book1.ISBN)
	ab1.Books = append(ab1.Books, book2.ISBN)

	var ab2 AuthorBooks
	ab2.Author = author2
	ab2.Books = append(ab2.Books, book1.ISBN)

	var ab3 AuthorBooks
	ab3.Author = author3
	ab3.Books = append(ab3.Books, book3.ISBN)

	AuthorList[SmStr(author1.Name)] = ab1
	AuthorList[SmStr(author2.Name)] = ab2
	AuthorList[SmStr(author3.Name)] = ab3

	BookList[book1.ISBN] = book1
	BookList[book2.ISBN] = book2
	BookList[book3.ISBN] = book3

	CredList[User.Username] = User.Password

	return
}

// Function signatures

func Login(w http.ResponseWriter, r *http.Request) {
	// Implement login logic
	w.Write([]byte("Hello World"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Implement logout logic
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get all books
}

func BookGeneralized(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get generalized book information
}

func GetSingleBook(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get a single book by ISBN
}

func NewBook(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create a new book
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Implement logic to delete a book by ISBN
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Implement logic to update a book by ISBN
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get all authors
}

func GetSingleAuthor(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get a single author by name
}

// SetupRouter initializes the router with all routes and middleware
func SetupRouter() chi.Router {
	// Create a new router
	r := chi.NewRouter()

	// Add middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	// Define the routes
	r.Post("/login", Login)
	r.Post("/logout", Logout)

	// Group routes for books and authors
	r.Group(func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Get("/", GetBooks)
			r.Get("/general", BookGeneralized)
			r.Get("/get/{ISBN}", GetSingleBook)

			// Group routes that require authentication
			r.Group(func(r chi.Router) {

				r.Post("/", NewBook)
				r.Delete("/{ISBN}", DeleteBook)
				r.Put("/{ISBN}", UpdateBook)
			})
		})

		r.Route("/authors", func(r chi.Router) {
			r.Get("/", GetAuthors)
			r.Get("/{AuthorName}", GetSingleAuthor)
		})
	})

	return r
}

func main() {
	IntializeDB()
	fmt.Println(AuthorList)
	fmt.Println(BookList)
	fmt.Println(CredList)
	// Setup the router
	r := SetupRouter()

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}
