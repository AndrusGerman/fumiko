package ports

import "github.com/AndrusGerman/fumiko/internal/core/domain"

type LLM interface {
	BasicQuest(text string) string
}

type LLMContext interface {
	GetMessages(id string) []*domain.Message
}
