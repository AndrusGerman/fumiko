package ports

type Config interface {
	GetTelegramToken() string
	GetDiscordToken() string
	GetBaseLLMContext() string

	EnableTelegram() bool
	EnableWhatsapp() bool
	EnableDiscord() bool
}
