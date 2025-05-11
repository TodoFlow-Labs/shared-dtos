package logging

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger = zerolog.Logger

func New(level string) Logger {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	return zerolog.New(os.Stdout).
		Level(lvl).
		With().
		Timestamp().
		Logger()
}
