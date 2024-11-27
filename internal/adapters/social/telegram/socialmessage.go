package telegram

import (
	"context"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type socialMessage struct {
	ctx    context.Context
	b      *bot.Bot
	update *models.Update
}

// GetText implements ports.SocialMessage.
func (s *socialMessage) GetText() string {
	return s.update.Message.Text
}

// ReplyText implements ports.SocialMessage.
func (s *socialMessage) ReplyText(text string) {
	s.b.SendMessage(s.ctx, &bot.SendMessageParams{
		ChatID: s.update.Message.Chat.ID,
		Text:   text,
	})
}

func newSocialMessage(ctx context.Context, b *bot.Bot, update *models.Update) ports.SocialMessage {
	return &socialMessage{
		ctx:    ctx,
		b:      b,
		update: update,
	}
}
