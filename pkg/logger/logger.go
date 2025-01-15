package logger

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var (
	logDir = "logs"
)

func NewLogger() (*LoggerManager, error) {

	fileName := fmt.Sprintf("%s/logs-%s.log", logDir, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetReportCaller(true)
	logger.SetLevel(log.TraceLevel)
	logger.SetOutput(io.MultiWriter(file, os.Stdout))

	return &LoggerManager{logger: logger}, nil
}

type LogFields log.Fields

type LogInput struct {
	Msg    interface{}
	Params LogFields
}

type LoggerManager struct {
	logger *log.Logger
}

func (l LoggerManager) Info(args ...interface{}) {
	info(l.logger, prepare(args...))
}

func (l LoggerManager) Debug(args ...interface{}) {
	debug(l.logger, prepare(args...))
}

func (l LoggerManager) Error(args ...interface{}) {
	err(l.logger, prepare(args...))
}

func (l LoggerManager) Warn(args ...interface{}) {
	warn(l.logger, prepare(args...))
}

func (l LoggerManager) Fatal(args ...interface{}) {
	fatal(l.logger, prepare(args...))
}

func (l LoggerManager) Panic(args ...interface{}) {
	panic(l.logger, prepare(args...))
}

func (l LoggerManager) PanicOnErr(err error, args ...interface{}) {
	if err != nil {
		l.Panic(prepare(args...))
	}
}

func (l LoggerManager) FatalOnErr(err error, args ...interface{}) {
	if err != nil {
		l.Panic(prepare(args...))
	}
}

func info(logger *log.Logger, input LogInput) {
	if len(input.Params) > 0 {
		logger.WithFields(log.Fields(input.Params)).Info(input.Msg)
	} else {
		logger.Info(input.Msg)
	}
}

func debug(logger *log.Logger, input LogInput) {
	if len(input.Params) > 0 {
		logger.WithFields(log.Fields(input.Params)).Debug(input.Msg)
	} else {
		logger.Debug(input.Msg)
	}
}

func warn(logger *log.Logger, input LogInput) {
	if len(input.Params) > 0 {
		logger.WithFields(log.Fields(input.Params)).Warn(input.Msg)
	} else {
		logger.Warn(input.Msg)
	}
}

func err(logger *log.Logger, input LogInput) {
	if len(input.Params) > 0 {
		logger.WithFields(log.Fields(input.Params)).Error(input.Msg)
	} else {
		logger.Error(input.Msg)
	}
}

func panic(logger *log.Logger, input LogInput) {
	if len(input.Params) > 0 {
		logger.WithFields(log.Fields(input.Params)).Panic(input.Msg)
	} else {
		logger.Panic(input.Msg)
	}
}

func fatal(logger *log.Logger, input LogInput) {
	if len(input.Params) > 0 {
		logger.WithFields(log.Fields(input.Params)).Fatal(input.Msg)
	} else {
		logger.Fatal(input.Msg)
	}
}

func prepare(args ...interface{}) LogInput {
	var fields = make(LogFields)
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
			fields["stacktrace"] = fmt.Sprintf("%+v", errors.WithStack(argErr))
		}
	}
	if strMsg != "" && errMsg != "" {
		return LogInput{Msg: fmt.Sprintf("%v: %v", strMsg, errMsg), Params: fields}
	} else if errMsg != "" {
		return LogInput{Msg: errMsg, Params: fields}
	}

	return LogInput{Msg: strMsg, Params: fields}
}
