package logging

type Logger interface {
	Log(level Level, msg string)
	Logf(level Level, msgFormat string, params ...any)
}
