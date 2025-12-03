package infra

import (
	"context"
	"net/http"

	"github.com/samEscom/BookHub/services/users/internal/handlers"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	mux      *http.ServeMux
	handlers *handlers.Handlers
	logger   *zap.Logger
	addr     string
}

func NewServer(mux *http.ServeMux, h *handlers.Handlers, logger *zap.Logger) *Server {
	s := &Server{
		mux:      mux,
		handlers: h,
		logger:   logger,
		addr:     ":8081",
	}

	s.routes()

	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/health", s.handlers.HealthCheck)
	s.mux.HandleFunc("/users", s.handlers.CreateUser)
	s.mux.HandleFunc("/users/", s.handlers.GetUserByID)
}

func StartServer(lc fx.Lifecycle, s *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				s.logger.Info("Users service running on", zap.String("addr", s.addr))
				if err := http.ListenAndServe(s.addr, s.mux); err != nil {
					s.logger.Error("Failed to start server", zap.Error(err))
				}
			}()
			return nil
		},
	})
}
