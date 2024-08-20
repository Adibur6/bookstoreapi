package apihandler

import (
	"fmt"
	"github.com/adibur6/bookstoreapi/authhandler"
	"github.com/adibur6/bookstoreapi/datahandler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
)

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
	r.Post("/login", authhandler.Login)
	r.Post("/logout", authhandler.Logout)

	// Group routes for books and authors
	r.Group(func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Use(jwtauth.Verifier(authhandler.TokenAuth))
			r.Use(jwtauth.Authenticator(authhandler.TokenAuth))
			r.Get("/", GetBooks)
			r.Get("/general", BookGeneralized)
			r.Get("/{ISBN}", GetSingleBook)

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
func Start(port int) {
	datahandler.InitializeDB()
	authhandler.InitToken()
	fmt.Println(datahandler.AuthorList)
	fmt.Println(datahandler.BookList)
	fmt.Println(datahandler.CredList)
	// Set up the router
	r := SetupRouter()

	// Start the HTTP server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Fatalln(err)
	}
}
