package services

import (
	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type FumikoService struct {
	llm        ports.LLM
	llmContext ports.LLMContext
}

// Quest implements ports.FumikoService.
func (f *FumikoService) Quest(userID domain.UserID, text string) (string, error) {
	// get llm last context
	var base = f.llmContext.GetMessages(userID)

	// get llm response
	var response, err = f.llm.Quest(base, text)
	if err != nil {
		return "", err
	}

	// add new message to context
	f.llmContext.AddMessages(userID, []*domain.Message{response})

	return response.Content, nil
}

func NewFumikoService(llm ports.LLM, llmContext ports.LLMContext) ports.FumikoService {
	return &FumikoService{
		llm:        llm,
		llmContext: llmContext,
	}
}
