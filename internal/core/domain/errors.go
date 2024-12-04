package domain

import (
	"errors"
)

// database
var ErrFailedOpenDatabase = errors.New("failed to open database")

// config
var ErrConfigTelegramTokenIsUndefined = errors.New("telgram token is undefined")
var ErrConfigDiscordTokenIsUndefined = errors.New("discord token is undefined")
