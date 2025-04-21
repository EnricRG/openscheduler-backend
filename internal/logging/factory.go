package logging

type LoggerFactory interface {
	GetLogger(params InitParams) Logger
}
