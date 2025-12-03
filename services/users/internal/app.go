package internal

import (
	"github.com/samEscom/BookHub/common/logger"
	"github.com/samEscom/BookHub/common/middleware"
	"github.com/samEscom/BookHub/common/mux"
	"github.com/samEscom/BookHub/services/users/internal/handlers"
	"github.com/samEscom/BookHub/services/users/internal/infra"
	"go.uber.org/fx"
)

func App() *fx.App {

	app := fx.New(
		fx.Provide(
			mux.NewMux,
			logger.NewLogger,
			middleware.NewLoggingMiddleware,
			handlers.NewHandlers,
			infra.NewServer,
		),
		fx.Invoke(func(*infra.Server) {}),
		fx.Invoke(infra.StartServer),
	)

	return app
}
