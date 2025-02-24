package interfaces

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Debug(args ...interface{})

	WithFields(args map[string]interface{}) Logger

	Errorf(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})

	PanicOnErr(err error, args ...interface{})
	FatalOnErr(err error, args ...interface{})
	OnErrorf(err error, msg string)
	OnError(err error, msg string)
}
