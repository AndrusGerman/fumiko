package ports

type Config interface {
	GetTelegramToken() string
	GetDiscordToken() string
	GetBaseLLMContext() string
}
