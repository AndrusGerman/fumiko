package ports

import "github.com/AndrusGerman/fumiko/internal/core/domain"

type LLM interface {
	BasicQuest(text string) (string, error)
	Quest(base []*domain.Message, text string) (*domain.Message, error)
	QuestParts(base []*domain.Message, text string, partsSize int) (<-chan *domain.Message, error)
}

type LLMContext interface {
	GetMessages(id domain.UserID) []*domain.Message
	AddMessages(id domain.UserID, messages []*domain.Message)
	SetMessages(id domain.UserID, messages []*domain.Message)
	ClearMessage(id domain.UserID)
}
