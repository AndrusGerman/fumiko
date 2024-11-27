package domain

import (
	"errors"
)

// database
var ErrFailedOpenDatabase = errors.New("failed to open database")

// config
var ErrConfigTelegramTokenIsUndefined = errors.New("telgram token is undefined")
