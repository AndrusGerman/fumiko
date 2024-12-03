package telegram

import (
	"context"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/NicoNex/echotron/v3"
)

type socialMessage struct {
	ctx    context.Context
	update *echotron.Update
	e      echotron.API
}

// GetText implements ports.SocialMessage.
func (s *socialMessage) GetText() string {
	return s.update.Message.Text
}

// ReplyText implements ports.SocialMessage.
func (s *socialMessage) ReplyText(text string) {
	s.e.SendMessage(text, s.update.ChatID(), &echotron.MessageOptions{})
}

func newSocialMessage(ctx context.Context, update *echotron.Update, e echotron.API) ports.SocialMessage {
	return &socialMessage{
		ctx:    ctx,
		update: update,
		e:      e,
	}
}
