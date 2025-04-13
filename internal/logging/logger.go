package logging

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func Setup(logFile *os.File) {
	logger = zerolog.New(logFile).With().Timestamp().Logger()
}

func GetLogger() zerolog.Logger {
	return logger
}
