package ports

import "github.com/AndrusGerman/fumiko/internal/core/domain"

type LLM interface {
	BasicQuest(text string) (string, error)
	Quest(base []*domain.Message, text string) (*domain.Message, error)
}

type LLMContext interface {
	GetMessages(id domain.UserID) []*domain.Message
	AddMessages(id domain.UserID, messages []*domain.Message)
}
