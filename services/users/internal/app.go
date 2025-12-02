package internal

import (
	"github.com/samEscom/BookHub/common/mux"
	"github.com/samEscom/BookHub/services/users/internal/handlers"
	"github.com/samEscom/BookHub/services/users/internal/infra"
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
	)

	return app
}
