package zerologger

import (
	"github.com/EnricRG/openscheduler-backend/internal/logging"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	Delegate zerolog.Logger //TODO make private
}

// // Log implements logging.Logger.
// func (z *ZeroLogger) Log(level logging.Level, msg string) {
// 	z.Delegate.WithLevel(mapLevel(level)).Msg(msg)
// }

// // Logf implements logging.Logger.
// func (z *ZeroLogger) Logf(level logging.Level, msgFormat string, params ...any) {
// 	z.Delegate.WithLevel(mapLevel(level)).Msgf(msgFormat, params...)
// }

func (z *ZeroLogger) Log(level logging.Level) {
	z.Delegate.WithLevel(mapLevel(level)).Send()
}

func (z *ZeroLogger) Msg(msg string) logging.LogEvent {
	z.Delegate.With().Str("message", msg)
}

func (z *ZeroLogger) Msgf(level logging.Level, msgFormat string, params ...any) {
	panic("not implemented") // TODO: Implement
}

func (z *ZeroLogger) Bool(key string, value bool) {
	panic("not implemented") // TODO: Implement
}

func (z *ZeroLogger) Str(key string, value string) {
	panic("not implemented") // TODO: Implement
}

func (z *ZeroLogger) Int64(key string, value int64) {
	panic("not implemented") // TODO: Implement
}

func (z *ZeroLogger) Float64(key string, value float64) {
	panic("not implemented") // TODO: Implement
}

func (z *ZeroLogger) Any(key string, value any) {
	panic("not implemented") // TODO: Implement
}

type factory struct {
}

// GetLogger implements logging.LoggerFactory.
func (f *factory) GetLogger(params logging.InitParams) logging.LogEvent {
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
