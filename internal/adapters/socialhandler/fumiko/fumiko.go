package fumiko

import "github.com/AndrusGerman/fumiko/internal/core/ports"

func NewFumikoHandler() ports.SocialHandler {
	return &FumikoHandler{}
}

type FumikoHandler struct {
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
	sm.ReplyText("Hola")
}
