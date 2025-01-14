package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func NewLogger() *LoggerManager {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)

	fileName := fmt.Sprintf("logs/logs-%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Warn("Не удалось открыть файл логов, используется вывод в консоль")
	}

	return &LoggerManager{}
}

type LogFields log.Fields

type LogInput struct {
	Msg    interface{}
	Params LogFields
}

type LoggerManager struct{}

func (l LoggerManager) Info(args ...interface{}) {
	Info(prepare(args...))
}

func (l LoggerManager) Debug(args ...interface{}) {
	Debug(prepare(args...))
}

func (l LoggerManager) Error(args ...interface{}) {
	Error(prepare(args...))
}

func (l LoggerManager) Warn(args ...interface{}) {
	Warn(prepare(args...))
}

func (l LoggerManager) Panic(args ...interface{}) {
	Panic(prepare(args...))
}

func (l LoggerManager) PanicOnFailed(err error, args ...interface{}) {
	if err != nil {
		l.Panic(prepare(args...))
	}
}

func Info(input LogInput) {
	if len(input.Params) > 0 {
		log.WithFields(log.Fields(input.Params)).Info(input.Msg)
	} else {
		log.Info(input.Msg)
	}
}

func Debug(input LogInput) {
	if len(input.Params) > 0 {
		log.WithFields(log.Fields(input.Params)).Debug(input.Msg)
	} else {
		log.Debug(input.Msg)
	}
}

func Warn(input LogInput) {
	if len(input.Params) > 0 {
		log.WithFields(log.Fields(input.Params)).Warn(input.Msg)
	} else {
		log.Warn(input.Msg)
	}
}

func Error(input LogInput) {
	if len(input.Params) > 0 {
		log.WithFields(log.Fields(input.Params)).Error(input.Msg)
	} else {
		log.Error(input.Msg)
	}
}

func Panic(input LogInput) {
	if len(input.Params) > 0 {
		log.WithFields(log.Fields(input.Params)).Panic(input.Msg)
	} else {
		log.Panic(input.Msg)
	}
}

func prepare(args ...interface{}) LogInput {
	var fields LogFields
	var strMsg string
	var errMsg string

	for _, arg := range args {
		if argFields, ok := arg.(map[string]interface{}); ok {
			fields = argFields
		}

		if argString, ok := arg.(string); ok {
			strMsg = argString
		}

		if argErr, ok := arg.(error); ok {
			errMsg = argErr.Error()
		}
	}

	if strMsg != "" && errMsg != "" {
		return LogInput{Msg: fmt.Sprintf("%v: %v", strMsg, errMsg), Params: fields}
	} else if errMsg != "" {
		return LogInput{Msg: errMsg, Params: fields}
	}

	return LogInput{Msg: strMsg, Params: fields}
}
