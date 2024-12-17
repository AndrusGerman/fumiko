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

	var responseStream, err = f.fumikoService.QuestParts(sm.GetUserID(), sm.GetText(), 140)
	if err != nil {
		sm.ReplyText("Fumiko😖: Error " + err.Error())
		return
	}

	for response := range responseStream {
		sm.ReplyText(response)
	}

}
