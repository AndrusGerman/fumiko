package discord

import (
	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/bwmarrin/discordgo"
)

type socialMessage struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

// GetUserName implements ports.SocialMessage.
func (s *socialMessage) GetUserName() string {
	return s.m.Author.Username
}

// GetUserID implements ports.SocialMessage.
func (s *socialMessage) GetUserID() domain.UserID {
	return domain.NewUserID(domain.DiscordSocialID, s.m.Author.Username)
}

// GetText implements ports.SocialMessage.
func (s *socialMessage) GetText() string {
	return s.m.Content
}

// ReplyText implements ports.SocialMessage.
func (s *socialMessage) ReplyText(text string) {
	s.s.ChannelMessageSend(s.m.ChannelID, text)
}

func newSocialMessage(s *discordgo.Session, m *discordgo.MessageCreate) ports.SocialMessage {
	return &socialMessage{
		s: s,
		m: m,
	}
}
