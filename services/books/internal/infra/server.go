package infra

import (
	"context"
	"log"
	"net/http"

	"github.com/samEscom/BookHub/services/books/internal/handlers"
	"go.uber.org/fx"
)

type Server struct {
	mux      *http.ServeMux
	handlers *handlers.Handlers
	addr     string
}

func NewServer(mux *http.ServeMux, h *handlers.Handlers) *Server {
	s := &Server{
		mux:      mux,
		handlers: h,
		addr:     ":8082",
	}

	s.routes()

	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/health", s.handlers.HealthCheck)
	s.mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		// Route to ListBooks for GET /books or CreateBook for POST /books
		if r.Method == http.MethodGet && r.URL.Path == "/books" {
			s.handlers.ListBooks(w, r)
		} else if r.Method == http.MethodPost {
			s.handlers.CreateBook(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	s.mux.HandleFunc("/books/", s.handlers.GetBookByID)
}

func StartServer(lc fx.Lifecycle, s *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Books service running on", s.addr)
				if err := http.ListenAndServe(s.addr, s.mux); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
	})
}
