package interfaces

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})

	PanicOnFailed(err error, args ...interface{})
}
