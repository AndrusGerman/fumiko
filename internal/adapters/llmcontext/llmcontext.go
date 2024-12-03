package llmcontext

import (
	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type LLMContext struct {
	messagesMemory map[string][]*domain.Message
}

// AddMessages implements ports.LLMContext.
func (lc *LLMContext) AddMessages(id domain.UserID, messages []*domain.Message) {
	lc.messagesMemory[id.String()] = append(lc.messagesMemory[id.String()], messages...)
}

func New() ports.LLMContext {
	return &LLMContext{
		messagesMemory: make(map[string][]*domain.Message),
	}
}

func (lc *LLMContext) GetMessages(id domain.UserID) []*domain.Message {
	return lc.messagesMemory[id.String()]
}
