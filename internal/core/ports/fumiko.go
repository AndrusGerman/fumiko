package ports

import "github.com/AndrusGerman/fumiko/internal/core/domain"

type FumikoService interface {
	Quest(userID domain.UserID, text string) (string, error)
	QuestParts(userID domain.UserID, text string, partSize int) (<-chan string, error)
}
