package logging

// Logger provides an implementation agnostic facade to emit logs. It must support parametric logging for LogEvent.
type Logger interface {
	Log(level Level, msg string)
	Logf(level Level, msgF string, args ...any)
	With(level Level) LogEvent
}

// LogEvent is the interface for implementation-agnostic domain level logging events. This interface encourages using structured
// loggers like Zerolog, but it's not a requirement.
//
// This interface requires that the implementations support context to allow rich log building. Logging events must not
// protect themselves from being emitted more than once.
//
// Implementations must also support the context modification methods that enable logging parameters. As previously stated,
// structured representation is not mandatory, but it has to represent each distinct parameter somehow. Duplicated keys should
// not be supported, so each implementation can choose how to deal with those situations.
//
// Once the event is logged, implementations can choose how to deal with potential repeated Log calls. It's recommended that,
// if possible, log events are re-logged and no panics occur.
type LogEvent interface {
	// Log is the final method of this event instance. If this method finishes successfully, a log request is expected to be
	// emitted to the underlying logger.
	Log()

	/// Context modification methods.

	Msg(msg string) LogEvent
	Msgf(msgFormat string, params ...any) LogEvent
	Bool(key string, value bool) LogEvent
	Str(key string, value string) LogEvent
	Int64(key string, value int64) LogEvent
	Float64(key string, value float64) LogEvent
	Any(key string, value any) LogEvent
	Err(err error) LogEvent
}
