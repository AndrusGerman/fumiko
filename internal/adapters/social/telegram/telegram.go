package telegram

import (
	"context"
	"fmt"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/fx"
)

type telegram struct {
	config         ports.Config
	socialHandlers []ports.SocialHandler
	b              *bot.Bot
}

// AddHandlers implements ports.Social.
func (t *telegram) AddHandlers(handlers ...ports.SocialHandler) {
	t.socialHandlers = handlers

}

// Register implements ports.Social.
func (t *telegram) Register() error {
	return nil
}

func (t *telegram) defaulHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update.Message.Text == "" {
		return
	}

	var socialMessage = newSocialMessage(ctx, b, update)
	fmt.Println("Received a telegram message!", socialMessage.GetText())

	for i := range t.socialHandlers {
		if t.socialHandlers[i].IsValid(socialMessage) {
			go t.socialHandlers[i].Message(socialMessage)
		}
	}

}

func (t *telegram) Start(c context.Context) error {
	opts := []bot.Option{
		bot.WithDefaultHandler(t.defaulHandler),
	}

	b, err := bot.New(t.config.GetTelegramToken(), opts...)
	if err != nil {
		return err
	}
	t.b = b

	go b.Start(c)

	return nil
}

func New(lc fx.Lifecycle, config ports.Config) ports.Social {
	var telegram = new(telegram)
	telegram.config = config

	lc.Append(fx.StartHook(telegram.Start))
	return telegram
}
