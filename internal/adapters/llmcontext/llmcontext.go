package llmcontext

import (
	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type LLMContext struct {
	messagesMemory map[string][]*domain.Message
}

func NewLLMContext() ports.LLMContext {
	return &LLMContext{
		messagesMemory: make(map[string][]*domain.Message),
	}
}

func (lc *LLMContext) GetMessages(id string) []*domain.Message {
	return lc.messagesMemory[id]
}
