package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/artkoro/hobbypal/userservice/internal/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	// TODO: Implement user creation logic

	// Return a success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get user logic
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update user logic
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete user logic
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement user registration logic
	// Example: Create a new user in the database
	// Extract user data from the request body
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	// TODO: Implement user registration logic (e.g., save user to the database)
	// After registration, you can return a success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User registered successfully")
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement user login logic
	// Example: Verify user credentials and generate a token
	// Extract user credentials from the request body
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	// TODO: Implement user login logic (e.g., verify credentials and generate a token)
	// After successful login, you can return the token or a success response
	token := "your-generated-token"
	response := map[string]string{
		"token": token,
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement user authentication logic
		// Example: Check if the request contains a valid token
		// Extract the token from the request headers or query parameters
		token := r.Header.Get("Authorization")
		// TODO: Validate the token (e.g., check if it's valid and belongs to an authenticated user)
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
