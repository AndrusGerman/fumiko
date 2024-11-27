package telegram

import (
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	tele "gopkg.in/telebot.v4"
)

type socialMessage struct {
	ctx tele.Context
	b   *tele.Bot
}

// GetText implements ports.SocialMessage.
func (s *socialMessage) GetText() string {
	return s.ctx.Text()
}

// ReplyText implements ports.SocialMessage.
func (s *socialMessage) ReplyText(text string) {
	s.ctx.Send(text)
}

func newSocialMessage(ctx tele.Context, b *tele.Bot) ports.SocialMessage {
	return &socialMessage{
		ctx: ctx,
		b:   b,
	}
}
