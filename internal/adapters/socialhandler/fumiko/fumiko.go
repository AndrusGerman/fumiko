package fumiko

import "github.com/AndrusGerman/fumiko/internal/core/ports"

func NewFumikoHandler(llm ports.LLM) ports.SocialHandler {
	return &FumikoHandler{
		llm: llm,
	}
}

type FumikoHandler struct {
	llm ports.LLM
}

// IsValid implements ports.SocialHandler.
func (f *FumikoHandler) IsValid(sm ports.SocialMessage) bool {
	return true
}

// Message implements ports.SocialHandler.
func (f *FumikoHandler) Message(sm ports.SocialMessage) {
	sm.GetText()
	var response = f.llm.BasicQuest(sm.GetText())

	sm.ReplyText("Fumiko💚:\n" + response)
}
