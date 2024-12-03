package fumiko

import "github.com/AndrusGerman/fumiko/internal/core/ports"

func NewFumikoHandler(llm ports.LLM, llmContext ports.LLMContext) ports.SocialHandler {
	return &FumikoHandler{
		llm:        llm,
		llmContext: llmContext,
	}
}

type FumikoHandler struct {
	llm        ports.LLM
	llmContext ports.LLMContext
}

// IsValid implements ports.SocialHandler.
func (f *FumikoHandler) IsValid(sm ports.SocialMessage) bool {
	return true
}

// Message implements ports.SocialHandler.
func (f *FumikoHandler) Message(sm ports.SocialMessage) {

	var base = f.llmContext.GetMessages(sm.GetUserID())
	var response, err = f.llm.Quest(base, sm.GetText())
	if err != nil {
		return
	}

	sm.ReplyText("Fumiko💚:\n" + response.Content)
}
