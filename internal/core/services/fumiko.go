package services

import (
	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type FumikoService struct {
	llm        ports.LLM
	llmContext ports.LLMContext
}

// QuestParts implements ports.FumikoService.
func (f *FumikoService) QuestParts(userID domain.UserID, text string, partSize int) (<-chan string, error) {

	var textStream = make(chan string)
	// get llm last context
	var base = f.llmContext.GetMessages(userID)

	// get llm response
	var responseStream, err = f.llm.QuestParts(base, text, partSize)
	if err != nil {
		return nil, err
	}

	go func() {
		var fullText string
		for response := range responseStream {
			fullText += response.Content
			textStream <- response.Content
		}
		// add new message to context
		f.llmContext.AddMessages(userID, []*domain.Message{domain.NewMessage(fullText, domain.AssistantRoleID)})

		close(textStream)
	}()

	return textStream, nil
	//return response.Content, nil
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
