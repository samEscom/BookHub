package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/samEscom/BookHub/services/users/internal/handlers/dto/request"
	"github.com/samEscom/BookHub/services/users/internal/handlers/dto/response"
	"go.uber.org/zap"
)

// CreateUser handles POST /users - creates a new user (mock)
func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req request.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Mock response
	mockUser := response.User{
		ID:        "usr_123456789",
		Username:  req.Username,
		Name:      req.Name,
		LastName:  req.LastName,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
		IsActive:  true,
	}

	h.logger.Info("User created", zap.String("username", req.Username))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mockUser)
}

// GetUserByID handles GET /users/{id} - retrieves a user by ID (mock)
func (h *Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	// For now, using a simple path parsing since we're using standard mux
	// The ID would be after /users/
	path := r.URL.Path
	id := path[len("/users/"):]

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Mock response
	mockUser := response.User{
		ID:        id,
		Username:  "john_doe",
		Name:      "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now().Add(-24 * time.Hour), // Created 1 day ago
		UpdatedAt: time.Now(),
		UpdatedBy: "admin",
		IsActive:  true,
	}

	h.logger.Info("User retrieved", zap.String("id", id))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockUser)
}
