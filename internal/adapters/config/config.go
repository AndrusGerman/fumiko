package config

import (
	"os"

	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/joho/godotenv"
)

type config struct {
	telegramToken string
}

// GetTelegramToken implements ports.Config.
func (c *config) GetTelegramToken() string {
	return c.telegramToken
}

func (c *config) Error() error {
	if c.telegramToken == "" {
		return domain.ErrConfigTelegramTokenIsUndefined
	}
	return nil
}

func New() (ports.Config, error) {
	godotenv.Load()
	var err error

	var config = &config{
		telegramToken: os.Getenv("TELEGRAM_TOKEN"),
	}

	if err = config.Error(); err != nil {
		return nil, err
	}

	return config, nil
}
