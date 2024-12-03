package ports

import "github.com/AndrusGerman/fumiko/internal/core/domain"

type Social interface {
	Register() error
	AddHandlers(handlers ...SocialHandler)
}

type SocialHandler interface {
	IsValid(sm SocialMessage) bool
	Message(sm SocialMessage)
}

type SocialMessage interface {
	GetText() string
	ReplyText(text string)
	GetUserID() domain.UserID
	GetUserName() string
}
