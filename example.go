package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

// UserService handles user-related operations
type UserService struct {
	users  []User
	nextID int
}

// NewUserService creates a new UserService instance
func NewUserService() *UserService {
	return &UserService{
		users:  make([]User, 0),
		nextID: 1,
	}
}

// CreateUser creates a new user and adds it to the service
func (s *UserService) CreateUser(name, email string) *User {
	user := User{
		ID:        s.nextID,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		Active:    true,
	}
	s.users = append(s.users, user)
	s.nextID++
	return &user
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (*User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// GetAllUsers returns all users in the system
func (s *UserService) GetAllUsers() []User {
	return s.users
}

// UpdateUser updates an existing user's information
func (s *UserService) UpdateUser(id int, name, email string) error {
	for i, user := range s.users {
		if user.ID == id {
			s.users[i].Name = name
			s.users[i].Email = email
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", id)
}

// DeactivateUser marks a user as inactive
func (s *UserService) DeactivateUser(id int) error {
	for i, user := range s.users {
		if user.ID == id {
			s.users[i].Active = false
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", id)
}

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	service *UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// ServeHTTP implements the http.Handler interface
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetUsers(w, r)
	case http.MethodPost:
		h.handleCreateUser(w, r)
	case http.MethodPut:
		h.handleUpdateUser(w, r)
	case http.MethodDelete:
		h.handleDeactivateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetUsers handles GET requests to retrieve users
func (h *UserHandler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	if idParam != "" {
		// Get specific user
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		user, err := h.service.GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	} else {
		// Get all users
		users := h.service.GetAllUsers()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// handleCreateUser handles POST requests to create a new user
func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	user := h.service.CreateUser(req.Name, req.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// handleUpdateUser handles PUT requests to update a user
func (h *UserHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateUser(id, req.Name, req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User updated successfully")
}

// handleDeactivateUser handles DELETE requests to deactivate a user
func (h *UserHandler) handleDeactivateUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeactivateUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User deactivated successfully")
}

// Example usage (this would normally be in a separate main package)
func exampleMain() {
	service := NewUserService()
	handler := NewUserHandler(service)

	// Create some sample users
	service.CreateUser("Alice Johnson", "alice@example.com")
	service.CreateUser("Bob Smith", "bob@example.com")
	service.CreateUser("Charlie Brown", "charlie@example.com")

	// Setup HTTP server
	http.Handle("/users", handler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
