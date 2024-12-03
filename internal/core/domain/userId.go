package domain

type SocialID string

const (
	WhatsappSocialID SocialID = "whatsapp"
	TelegramSocialID SocialID = "telegram"
)

type UserID struct {
	socialID SocialID
	id       string
}

func (u *UserID) SocialID() SocialID {
	return u.socialID
}

func (u *UserID) ID() string {
	return u.id
}

func NewUserID(socialID SocialID, id string) UserID {
	return UserID{socialID: socialID, id: id}
}
