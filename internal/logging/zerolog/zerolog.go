package zerologger

import (
	"github.com/EnricRG/openscheduler-backend/internal/logging"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	Delegate zerolog.Logger
}

// Log implements logging.Logger.
func (z *ZeroLogger) Log(level logging.Level, msg string) {
	z.Delegate.WithLevel(mapLevel(level)).Msg(msg)
}

// Logf implements logging.Logger.
func (z *ZeroLogger) Logf(level logging.Level, msgFormat string, params ...any) {
	z.Delegate.WithLevel(mapLevel(level)).Msgf(msgFormat, params...)
}

type factory struct {
}

// GetLogger implements logging.LoggerFactory.
func (f *factory) GetLogger(params logging.InitParams) logging.Logger {
	return NewLogger(params)
}

func NewLogger(params logging.InitParams) *ZeroLogger {
	return &ZeroLogger{
		Delegate: zerolog.New(params.Writer).Level(mapLevel(params.Level)).With().Timestamp().Logger(),
	}
}

func NewFactory() logging.LoggerFactory {
	return &factory{}
}

func mapLevel(level logging.Level) (zerolevel zerolog.Level) {
	switch level {
	case logging.Trace:
		zerolevel = zerolog.TraceLevel
	case logging.Debug:
		zerolevel = zerolog.DebugLevel
	case logging.Info:
		zerolevel = zerolog.InfoLevel
	case logging.Warn:
		zerolevel = zerolog.WarnLevel
	case logging.Error:
		zerolevel = zerolog.ErrorLevel
	}
	return
}
