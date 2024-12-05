package dump

import (
	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
)

type dump struct {
}

// GetSocialID implements ports.Social.
func (d *dump) GetSocialID() domain.SocialID {
	return domain.DumpSocialID
}

// AddHandlers implements ports.Social.
func (d *dump) AddHandlers(handlers ...ports.SocialHandler) {

}

// Register implements ports.Social.
func (d *dump) Register() error {
	return nil
}

func New() ports.Social {
	return new(dump)
}
