package bootstrap

import (
	"github.com/AndrusGerman/fumiko/internal/adapters/social/whatsapp"
	"go.uber.org/fx"
)

// deps
func privide() fx.Option {
	return fx.Provide(
		// social manager
		whatsapp.New,
	)
}

func Run() {
	var app = fx.New(
		privide(),
		fx.Invoke(start),
	)

	app.Run()
}

func start() {

}
