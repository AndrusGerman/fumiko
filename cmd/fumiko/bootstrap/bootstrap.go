package bootstrap

import (
	"fmt"

	"github.com/AndrusGerman/fumiko/internal/adapters/llm/ollama"
	"github.com/AndrusGerman/fumiko/internal/adapters/social/whatsapp"
	"github.com/AndrusGerman/fumiko/internal/adapters/socialhandler"
	"github.com/AndrusGerman/fumiko/internal/adapters/socialhandler/fumiko"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"go.uber.org/fx"
)

// deps
func privide() fx.Option {
	return fx.Provide(
		// social manager
		whatsapp.New,

		// llm
		ollama.New,
	)
}

func socialhandlerProvide() fx.Option {
	return socialhandler.NewHandlers(
		fumiko.NewFumikoHandler,
	)
}

func Run() {
	var app = fx.New(
		privide(),
		socialhandlerProvide(),
		socialhandler.NewHandlers(
			fumiko.NewFumikoHandler,
		),
		fx.Invoke(start),
	)

	app.Run()
}

type startDto struct {
	fx.In
	Social         ports.Social
	SocialHandlers []ports.SocialHandler `group:"socialHandlers"`
}

func start(dto startDto) {
	dto.Social.AddHandlers(dto.SocialHandlers...)
	fmt.Println("Handlers: ", dto.SocialHandlers)
}
