package llmcontext

import (
	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type LLMContext struct {
	messagesMemory map[string][]*domain.Message
	config         ports.Config
}

// ClearMessage implements ports.LLMContext.
func (lc *LLMContext) ClearMessage(id domain.UserID) {
	lc.messagesMemory[id.String()] = []*domain.Message{}
}

// SetMessages implements ports.LLMContext.
func (lc *LLMContext) SetMessages(id domain.UserID, messages []*domain.Message) {
	lc.messagesMemory[id.String()] = messages
}

// AddMessages implements ports.LLMContext.
func (lc *LLMContext) AddMessages(id domain.UserID, messages []*domain.Message) {
	lc.messagesMemory[id.String()] = append(lc.messagesMemory[id.String()], messages...)
}

func New(config ports.Config) ports.LLMContext {
	return &LLMContext{
		messagesMemory: make(map[string][]*domain.Message),
		config:         config,
	}
}

func (lc *LLMContext) GetMessages(id domain.UserID) []*domain.Message {
	if len(lc.messagesMemory[id.String()]) > 0 {
		return lc.messagesMemory[id.String()]
	}

	return []*domain.Message{domain.NewMessage(lc.config.GetBaseLLMContext(), domain.SystemRoleID)}
}
