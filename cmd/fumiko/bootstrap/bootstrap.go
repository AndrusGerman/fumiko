package bootstrap

import (
	"fmt"

	"github.com/AndrusGerman/fumiko/internal/adapters/config"
	"github.com/AndrusGerman/fumiko/internal/adapters/llm/ollama"
	"github.com/AndrusGerman/fumiko/internal/adapters/llmcontext"
	"github.com/AndrusGerman/fumiko/internal/adapters/rest"
	"github.com/AndrusGerman/fumiko/internal/adapters/social"
	"github.com/AndrusGerman/fumiko/internal/adapters/social/discord"
	"github.com/AndrusGerman/fumiko/internal/adapters/social/telegram"
	"github.com/AndrusGerman/fumiko/internal/adapters/social/whatsapp"
	"github.com/AndrusGerman/fumiko/internal/adapters/socialhandler"
	"github.com/AndrusGerman/fumiko/internal/adapters/socialhandler/fumiko"
	"github.com/AndrusGerman/fumiko/internal/adapters/storage/sqlite3"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/AndrusGerman/fumiko/internal/core/services"
	"go.uber.org/fx"
)

// core Deps
func coreDeps() fx.Option {
	return fx.Provide(
		//config
		config.New,

		// database
		sqlite3.New,

		// rest
		rest.New,

		// llm
		llmcontext.New,
		ollama.New,

		// services
		services.NewFumikoService,
	)
}

// social manager
func socials() fx.Option {
	return social.NewSocials(
		whatsapp.New,
		telegram.New,
		discord.New,
	)
}

// handlers
func socialhandlerProvide() fx.Option {
	return socialhandler.NewHandlers(
		fumiko.NewFumikoHandler,
	)
}

func Run() {
	var app = fx.New(
		coreDeps(),
		socialhandlerProvide(),
		socials(),
		fx.Invoke(start),
	)

	app.Run()
}

type startDto struct {
	fx.In
	Social         []ports.Social        `group:"social"`
	SocialHandlers []ports.SocialHandler `group:"socialHandlers"`
}

func start(dto startDto) {
	for i := range dto.Social {
		dto.Social[i].AddHandlers(dto.SocialHandlers...)
	}
	fmt.Println("Handlers: ", len(dto.SocialHandlers))
	fmt.Println("Socials: ", len(dto.Social))

}
