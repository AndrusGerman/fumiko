package telegram

import (
	"time"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	tele "gopkg.in/telebot.v4"
)

type telegram struct {
	b              *tele.Bot
	socialHandlers []ports.SocialHandler
}

// AddHandlers implements ports.Social.
func (t *telegram) AddHandlers(handlers ...ports.SocialHandler) {
	t.socialHandlers = handlers

}

// Register implements ports.Social.
func (t *telegram) Register() error {
	return nil
}

func (t *telegram) registerMessage() {
	t.b.Handle(tele.OnText, func(ctx tele.Context) error {

		var socialMessage = newSocialMessage(ctx, t.b)
		for i := range t.socialHandlers {
			if t.socialHandlers[i].IsValid(socialMessage) {
				t.socialHandlers[i].Message(socialMessage)
			}
		}

		return nil
	})

}

func New(config ports.Config) (ports.Social, error) {
	var telegram = new(telegram)
	var b *tele.Bot
	var err error

	pref := tele.Settings{
		Token:  config.GetTelegramToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	if b, err = tele.NewBot(pref); err != nil {
		return nil, err
	}
	telegram.b = b

	return telegram, nil
}
