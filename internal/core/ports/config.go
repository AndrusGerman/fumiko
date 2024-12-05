package ports

import "github.com/AndrusGerman/fumiko/internal/core/domain"

type Config interface {
	GetTelegramToken() string
	GetDiscordToken() string
	GetBaseLLMContext() string

	EnableSocial(socialID domain.SocialID) bool
}
