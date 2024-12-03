package fumiko

import "github.com/AndrusGerman/fumiko/internal/core/ports"

func NewFumikoHandler(fumikoService ports.FumikoService) ports.SocialHandler {
	return &FumikoHandler{
		fumikoService: fumikoService,
	}
}

type FumikoHandler struct {
	fumikoService ports.FumikoService
}

// IsValid implements ports.SocialHandler.
func (f *FumikoHandler) IsValid(sm ports.SocialMessage) bool {
	return true
}

// Message implements ports.SocialHandler.
func (f *FumikoHandler) Message(sm ports.SocialMessage) {

	var response, err = f.fumikoService.Quest(sm.GetUserID(), sm.GetText())
	if err != nil {
		sm.ReplyText("Fumiko😖: Error " + err.Error())
		return
	}
	sm.ReplyText("Fumiko💚:\n" + response)
}
