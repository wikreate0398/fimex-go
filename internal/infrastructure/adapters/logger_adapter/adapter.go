package logger_adapter

import (
	"github.com/sirupsen/logrus"
	"wikreate/fimex/internal/domain/interfaces"
)

var _ interfaces.Logger = (*LoggerAdapter)(nil)

type LogFields logrus.Fields

type LogInput struct {
	Msg    interface{}
	Params LogFields
}

type LoggerAdapter struct {
	logger *logrus.Logger
	fields LogFields
}

func NewLoggerAdapter(logger *logrus.Logger) *LoggerAdapter {
	return &LoggerAdapter{logger: logger}
}

func (l *LoggerAdapter) WithFields(args map[string]interface{}) interfaces.Logger {
	var fields = make(LogFields)

	for k, v := range args {
		fields[k] = v
	}

	l.fields = fields

	return l
}

func (l *LoggerAdapter) Info(args ...interface{}) {
	l.fillFields().Info(args...)
}

func (l *LoggerAdapter) Debug(args ...interface{}) {
	l.fillFields().Debug(args...)
}

func (l *LoggerAdapter) Debugf(msg string, args ...interface{}) {
	l.fillFields().Debugf(msg, args...)
}

func (l *LoggerAdapter) Error(args ...interface{}) {
	l.fillFields().Error(args...)
}

func (l *LoggerAdapter) Errorf(msg string, args ...interface{}) {
	l.fillFields().Errorf(msg, args...)
}

func (l *LoggerAdapter) OnErrorf(err error, msg string) {
	if err != nil {
		l.Errorf(msg, err)
	}
}

func (l *LoggerAdapter) OnError(err error, msg string) {
	if err != nil {
		l.Error(msg, err)
	}
}

func (l *LoggerAdapter) Warn(args ...interface{}) {
	l.fillFields().Warn(args...)
}

func (l *LoggerAdapter) Fatal(args ...interface{}) {
	l.fillFields().Fatal(args...)
}

func (l *LoggerAdapter) Panic(args ...interface{}) {
	l.fillFields().Panic(args...)
}

func (l *LoggerAdapter) PanicOnErr(err error, args ...interface{}) {
	if err != nil {
		l.Panic(append(args, err)...)
	}
}

func (l *LoggerAdapter) FatalOnErr(err error, args ...interface{}) {
	if err != nil {
		l.Fatal(append(args, err)...)
	}
}

func (l *LoggerAdapter) fillFields() *logrus.Entry {
	var instance = logrus.NewEntry(l.logger)

	if len(l.fields) > 0 {
		instance = instance.WithFields(logrus.Fields(l.fields))
	}

	l.fields = nil

	return instance
}
