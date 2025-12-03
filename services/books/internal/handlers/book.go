package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/samEscom/BookHub/services/books/internal/handlers/dto/request"
	"github.com/samEscom/BookHub/services/books/internal/handlers/dto/response"
)

// CreateBook handles POST /books - creates a new book (mock)
func (h *Handlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req request.CreateBook
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Mock response
	mockBook := response.Book{
		ID:            "book_123456789",
		Title:         req.Title,
		Author:        req.Author,
		ISBN:          req.ISBN,
		PublishedYear: req.PublishedYear,
		Genre:         req.Genre,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		IsAvailable:   true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mockBook)
}

// GetBookByID handles GET /books/{id} - retrieves a book by ID (mock)
func (h *Handlers) GetBookByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	path := r.URL.Path
	id := path[len("/books/"):]

	if id == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}

	// Mock response
	mockBook := response.Book{
		ID:            id,
		Title:         "The Go Programming Language",
		Author:        "Alan Donovan & Brian Kernighan",
		ISBN:          "978-0134190440",
		PublishedYear: 2015,
		Genre:         "Technology",
		CreatedAt:     time.Now().Add(-30 * 24 * time.Hour), // Created 30 days ago
		UpdatedAt:     time.Now(),
		IsAvailable:   true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockBook)
}

// ListBooks handles GET /books - lists all books (mock)
func (h *Handlers) ListBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mock response with sample books
	mockBooks := []response.Book{
		{
			ID:            "book_001",
			Title:         "The Go Programming Language",
			Author:        "Alan Donovan & Brian Kernighan",
			ISBN:          "978-0134190440",
			PublishedYear: 2015,
			Genre:         "Technology",
			CreatedAt:     time.Now().Add(-30 * 24 * time.Hour),
			UpdatedAt:     time.Now(),
			IsAvailable:   true,
		},
		{
			ID:            "book_002",
			Title:         "Clean Code",
			Author:        "Robert C. Martin",
			ISBN:          "978-0132350884",
			PublishedYear: 2008,
			Genre:         "Software Engineering",
			CreatedAt:     time.Now().Add(-20 * 24 * time.Hour),
			UpdatedAt:     time.Now(),
			IsAvailable:   true,
		},
		{
			ID:            "book_003",
			Title:         "Design Patterns",
			Author:        "Gang of Four",
			ISBN:          "978-0201633610",
			PublishedYear: 1994,
			Genre:         "Software Engineering",
			CreatedAt:     time.Now().Add(-10 * 24 * time.Hour),
			UpdatedAt:     time.Now(),
			IsAvailable:   false,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockBooks)
}
