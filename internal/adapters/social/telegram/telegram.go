package telegram

import (
	"context"
	"fmt"

	"github.com/AndrusGerman/fumiko/internal/adapters/social/dump"
	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/NicoNex/echotron/v3"

	"go.uber.org/fx"
)

type telegram struct {
	config         ports.Config
	socialHandlers []ports.SocialHandler
	e              echotron.API
}

// AddHandlers implements ports.Social.
func (t *telegram) AddHandlers(handlers ...ports.SocialHandler) {
	t.socialHandlers = handlers

}

// Register implements ports.Social.
func (t *telegram) Register() error {
	return nil
}

func (t *telegram) defaulHandler(update *echotron.Update) {
	if update.Message.Text == "" {
		return
	}

	var socialMessage = newSocialMessage(context.TODO(), update, t.e)
	fmt.Println("Received a telegram message!", socialMessage.GetText())

	for i := range t.socialHandlers {
		if t.socialHandlers[i].IsValid(socialMessage) {
			go t.socialHandlers[i].Message(socialMessage)
		}
	}

}

func (t *telegram) Start(c context.Context) error {
	t.e = echotron.NewAPI(t.config.GetTelegramToken())

	go func() {
		for update := range echotron.PollingUpdates(t.config.GetTelegramToken()) {
			t.defaulHandler(update)
		}
	}()

	return nil
}

// GetSocialID implements ports.Social.
func (d *telegram) GetSocialID() domain.SocialID {
	return domain.TelegramSocialID
}

func New(lc fx.Lifecycle, config ports.Config) ports.Social {
	if !config.EnableSocial(domain.TelegramSocialID) {
		return dump.New()
	}

	var telegram = new(telegram)
	telegram.config = config
	lc.Append(fx.StartHook(telegram.Start))
	return telegram
}
