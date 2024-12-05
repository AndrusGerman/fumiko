package config

import (
	"flag"
	"os"

	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/joho/godotenv"
)

type config struct {
	telegramToken  string
	discordToken   string
	baseLLMContext string

	enableDiscord  bool
	enableTelegram bool
	enableWhatsapp bool
}

// EnableDiscord implements ports.Config.
func (c *config) EnableDiscord() bool {
	return c.enableDiscord
}

// EnableTelegram implements ports.Config.
func (c *config) EnableTelegram() bool {
	return c.enableTelegram
}

// EnableWhatsapp implements ports.Config.
func (c *config) EnableWhatsapp() bool {
	return c.enableWhatsapp
}

// GetBaseLLMContext implements ports.Config.
func (c *config) GetBaseLLMContext() string {
	return c.baseLLMContext
}

// GetDiscordToken implements ports.Config.
func (c *config) GetDiscordToken() string {
	return c.discordToken
}

// GetTelegramToken implements ports.Config.
func (c *config) GetTelegramToken() string {
	return c.telegramToken
}

func (c *config) Error() error {
	if c.telegramToken == "" && c.enableTelegram {
		return domain.ErrConfigTelegramTokenIsUndefined
	}
	if c.discordToken == "" && c.enableDiscord {
		return domain.ErrConfigDiscordTokenIsUndefined
	}
	if c.baseLLMContext == "" {
		return domain.ErrConfigBaseLLMContextIsUndefined
	}
	return nil
}

func (c *config) ReadEnv() error {
	c.telegramToken = os.Getenv("TELEGRAM_TOKEN")
	c.discordToken = os.Getenv("DISCORD_TOKEN")
	c.baseLLMContext = os.Getenv("BASE_LLMCONTEXT")
	return nil
}

func (c *config) ReadFlags() error {
	flag.BoolVar(&c.enableDiscord, "discord", false, "enable or disable discord")
	flag.BoolVar(&c.enableTelegram, "telegram", false, "enable or disable telegram")
	flag.BoolVar(&c.enableWhatsapp, "whatsapp", false, "enable or disable whatsapp")

	flag.Parse()
	return nil
}

func New() (ports.Config, error) {
	godotenv.Load()
	var err error

	var config = new(config)

	if err = config.ReadEnv(); err != nil {
		return nil, err
	}

	if err = config.ReadFlags(); err != nil {
		return nil, err
	}

	if err = config.Error(); err != nil {
		return nil, err
	}

	return config, nil
}
