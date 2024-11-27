package whatsapp

import (
	"context"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
)

type socialMessage struct {
	event  *events.Message
	client *whatsmeow.Client
}

// GetText implements ports.SocialMessage.
func (s *socialMessage) GetText() string {
	return s.event.Message.GetConversation()[1:]
}

// ReplyText implements ports.SocialMessage.
func (s *socialMessage) ReplyText(text string) {
	s.client.SendMessage(context.Background(), s.event.Info.Chat, &waE2E.Message{
		Conversation: &text,
	})
}

func newSocialMessage(event *events.Message, client *whatsmeow.Client) ports.SocialMessage {
	return &socialMessage{
		event:  event,
		client: client,
	}
}
