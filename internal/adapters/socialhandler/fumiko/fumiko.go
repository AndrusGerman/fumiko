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
	var message = sm.GetText()
	if len(message) < 1 {
		return false
	}
	return string(message[0]) == "."
}

// Message implements ports.SocialHandler.
func (f *FumikoHandler) Message(sm ports.SocialMessage) {
	var response = f.llm.BasicQuest(sm.GetText()[1:])

	sm.ReplyText(response)

}
