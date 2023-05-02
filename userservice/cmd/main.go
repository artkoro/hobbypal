package main

import (
	"net/http"

	"github.com/artkoro/hobbypal/userservice/config"
	"github.com/artkoro/hobbypal/userservice/handlers"

	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement user registration logic
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement user login logic
}
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement user authentication logic

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Connect to the database
	db := config.Connect()
	defer db.Close()

	// Create a new router
	router := mux.NewRouter()

	// Define the routes and map them to the handler functions
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Protected routes that require authentication
	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(handlers.Authenticate)

	// Define the protected routes
	protectedRoutes.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	protectedRoutes.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Start the server
	http.ListenAndServe(":8000", router)
}
