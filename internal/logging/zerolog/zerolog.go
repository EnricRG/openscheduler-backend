// zerologger provides a logging package interfaces implementation using Zerolog.
package zerologger

import (
	"fmt"

	"github.com/EnricRG/openscheduler-backend/internal/logging"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	Delegate zerolog.Logger //TODO make private
}

// Log implements logging.Logger.
func (z *ZeroLogger) Log(level logging.Level, msg string) {
	z.Delegate.WithLevel(mapLevel(level)).Msg(msg)
}

// Logf implements logging.Logger.
func (z *ZeroLogger) Logf(level logging.Level, msgFormat string, params ...any) {
	z.Delegate.WithLevel(mapLevel(level)).Msgf(msgFormat, params...)
}

func (z *ZeroLogger) With(level logging.Level) logging.LogEvent {
	return &ZeroLogEvent{Delegate: z.Delegate.WithLevel(mapLevel(level))}
}

type ZeroLogEvent struct {
	Delegate *zerolog.Event
}

func (z *ZeroLogEvent) Log() {
	z.Delegate.Send()
}

func (z *ZeroLogEvent) Msg(msg string) logging.LogEvent {
	z.Delegate.Str("message", msg)
	return z
}

func (z *ZeroLogEvent) Msgf(msgFormat string, params ...any) logging.LogEvent {
	z.Delegate.Str("message", fmt.Sprintf(msgFormat, params...))
	return z
}

func (z *ZeroLogEvent) Bool(key string, value bool) logging.LogEvent {
	z.Delegate.Bool(key, value)
	return z
}

func (z *ZeroLogEvent) Str(key string, value string) logging.LogEvent {
	z.Delegate.Str(key, value)
	return z
}

func (z *ZeroLogEvent) Int64(key string, value int64) logging.LogEvent {
	z.Delegate.Int64(key, value)
	return z
}

func (z *ZeroLogEvent) Float64(key string, value float64) logging.LogEvent {
	z.Delegate.Float64(key, value)
	return z
}

func (z *ZeroLogEvent) Any(key string, value any) logging.LogEvent {
	z.Delegate.Any(key, value)
	return z
}

func (z *ZeroLogEvent) Err(err error) logging.LogEvent {
	z.Delegate.Err(err)
	return z
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
