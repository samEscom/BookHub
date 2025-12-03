package internal

import (
	"github.com/samEscom/BookHub/common/mux"
	"github.com/samEscom/BookHub/services/books/internal/handlers"
	"github.com/samEscom/BookHub/services/books/internal/infra"
	"go.uber.org/fx"
)

func App() *fx.App {

	app := fx.New(
		fx.Provide(
			mux.NewMux,
			handlers.NewHandlers,
			infra.NewServer,
		),
		fx.Invoke(func(*infra.Server) {}),
		fx.Invoke(infra.StartServer),
	)

	return app
}
