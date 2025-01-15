package interfaces

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})

	PanicOnErr(err error, args ...interface{})
	FatalOnErr(err error, args ...interface{})
}
