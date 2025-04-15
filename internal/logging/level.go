package logging

type Level uint8

const (
	// Trace defines trace log level.
	Trace Level = iota
	// Debug defines debug log level.
	Debug
	// Info defines info log level.
	Info
	// Warn defines warn log level.
	Warn
	// Error defines error log level.
	Error
)
