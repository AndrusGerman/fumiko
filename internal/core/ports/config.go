package ports

type Config interface {
	GetTelegramToken() string
}
