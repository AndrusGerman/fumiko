package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/AndrusGerman/fumiko/internal/core/domain"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"github.com/joho/godotenv"
)

type config struct {
	telegramToken  string
	discordToken   string
	baseLLMContext string

	flagsSocial map[domain.SocialID]*bool
}

// EnableSocial implements ports.Config.
func (c *config) EnableSocial(socialID domain.SocialID) bool {
	return *c.flagsSocial[socialID]
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
	if c.telegramToken == "" && c.EnableSocial(domain.TelegramSocialID) {
		return domain.ErrConfigTelegramTokenIsUndefined
	}

	if c.discordToken == "" && c.EnableSocial(domain.DiscordSocialID) {
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
	for _, social := range domain.AppSocialList {
		var socialValue = false
		c.flagsSocial[social] = &socialValue
		flag.BoolVar(c.flagsSocial[social], social.String(), false, fmt.Sprintf("enable or disable %s", social))
	}

	flag.Parse()
	return nil
}

func New() (ports.Config, error) {
	godotenv.Load()
	var err error

	var config = &config{
		flagsSocial: make(map[domain.SocialID]*bool),
	}

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
