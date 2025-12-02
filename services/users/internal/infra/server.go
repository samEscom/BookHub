package infra

import (
	"context"
	"log"
	"net/http"

	"github.com/samEscom/BookHub/services/users/internal/handlers"
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
		addr:     ":8081",
	}

	s.routes()

	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/health", s.handlers.HealthCheck)
}

func StartServer(lc fx.Lifecycle, s *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Users service running on", s.addr)
				if err := http.ListenAndServe(s.addr, s.mux); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
	})
}
