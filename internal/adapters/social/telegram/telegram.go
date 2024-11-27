package telegram

import (
	"fmt"
	"time"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"go.uber.org/fx"
	tele "gopkg.in/telebot.v4"
)

type telegram struct {
	b              *tele.Bot
	config         ports.Config
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
	// t.b.Handle("/start", func(ctx tele.Context) error {
	// 	fmt.Println("Start??")
	// 	return nil
	// })
	t.b.Handle(tele.OnText, func(ctx tele.Context) error {
		var socialMessage = newSocialMessage(ctx, t.b)
		fmt.Println("Received a telegram message!", socialMessage.GetText())

		for i := range t.socialHandlers {
			if t.socialHandlers[i].IsValid(socialMessage) {
				t.socialHandlers[i].Message(socialMessage)
			}
		}

		return nil
	})

}

func (t *telegram) Start() error {
	var b *tele.Bot
	var err error

	pref := tele.Settings{
		Token:  t.config.GetTelegramToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	if b, err = tele.NewBot(pref); err != nil {
		return err
	}
	t.b = b

	t.registerMessage()

	return nil
}

func New(lc fx.Lifecycle, config ports.Config) ports.Social {
	var telegram = new(telegram)
	telegram.config = config

	lc.Append(fx.StartHook(telegram.Start))
	return telegram
}
