package telegram

import "github.com/AndrusGerman/fumiko/internal/core/ports"

type telegram struct {
	socialHandlers []ports.SocialHandler
}

// AddHandlers implements ports.Social.
func (t *telegram) AddHandlers(handlers ...ports.SocialHandler) {
	t.socialHandlers = handlers

}

// Register implements ports.Social.
func (t *telegram) Register() error {
	return nil
}

func New(config ports.Config) ports.Social {
	var telegram = new(telegram)

	return telegram
}
