package whatsapp

import (
	"context"

	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
)

type socialMessage struct {
	event  *events.Message
	client *whatsmeow.Client
	w      *whatsapp
}

// GetUserName implements ports.SocialMessage.
func (s *socialMessage) GetUserName() string {
	return s.event.Info.PushName
}

// GetUserID implements ports.SocialMessage.
func (s *socialMessage) GetUserID() domain.UserID {
	return domain.NewUserID(domain.WhatsappSocialID, s.event.Info.Sender.ADString())
}

// GetText implements ports.SocialMessage.
func (s *socialMessage) GetText() string {
	var ignoreTextLength = len(s.w.keyText)
	return s.event.Message.GetConversation()[ignoreTextLength:]
}

// ReplyText implements ports.SocialMessage.
func (s *socialMessage) ReplyText(text string) {
	s.client.SendMessage(context.Background(), s.event.Info.Chat, &waE2E.Message{
		Conversation: &text,
	})
}

func newSocialMessage(event *events.Message, client *whatsmeow.Client, w *whatsapp) ports.SocialMessage {
	return &socialMessage{
		event:  event,
		client: client,
		w:      w,
	}
}
