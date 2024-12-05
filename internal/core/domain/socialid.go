package domain

type SocialID string

const (
	WhatsappSocialID SocialID = "whatsapp"
	TelegramSocialID SocialID = "telegram"
	DiscordSocialID  SocialID = "discord"

	DumpSocialID SocialID = "dump"
)

func (s SocialID) String() string {
	return string(s)
}

var AppSocialList = []SocialID{
	WhatsappSocialID,
	TelegramSocialID,
	DiscordSocialID,
}
