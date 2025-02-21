package logger

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

type LogFields log.Fields

type LogInput struct {
	Msg    interface{}
	Params LogFields
}

type LoggerManager struct {
	logger *log.Logger
	fields LogFields
}

func (l *LoggerManager) WithFields(args map[string]interface{}) *LoggerManager {
	var fields = make(LogFields)

	for k, v := range args {
		fields[k] = v
	}

	l.fields = fields

	return l
}

func (l *LoggerManager) Info(args ...interface{}) {
	info(l.logger, prepare(args...))
}

func (l *LoggerManager) Debug(args ...interface{}) {
	debug(l.logger, prepare(args...))
}

func (l *LoggerManager) Debugf(msg string, args ...interface{}) {
	debugf(l.logger, msg, prepare(args...))
}

func (l *LoggerManager) Error(args ...interface{}) {
	err(l.logger, prepare(args...))
}

func (l *LoggerManager) Errorf(msg string, args ...interface{}) {
	l.fillFields().Errorf(msg, args...)
}

func (l *LoggerManager) fillFields() *log.Entry {
	var instance = log.NewEntry(l.logger)

	if len(l.fields) > 0 {
		instance = instance.WithFields(log.Fields(l.fields))
	}

	fmt.Println(len(l.fields), "$$$")

	l.fields = nil

	return instance
}

func (l *LoggerManager) Warn(args ...interface{}) {
	warn(l.logger, prepare(args...))
}

func (l *LoggerManager) Fatal(args ...interface{}) {
	fatal(l.logger, prepare(args...))
}

func (l *LoggerManager) Panic(args ...interface{}) {
	ppanic(l.logger, prepare(args...))
}

func (l *LoggerManager) PanicOnErr(err error, args ...interface{}) {
	if err != nil {
		l.Panic(append(args, err)...)
	}
}

func (l *LoggerManager) FatalOnErr(err error, args ...interface{}) {
	if err != nil {
		l.Fatal(append(args, err)...)
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

func debugf(logger *log.Logger, msg string, input LogInput) {
	if len(input.Params) > 0 {
		logger.WithFields(log.Fields(input.Params)).Debug(msg)
	} else {
		logger.Debugf(msg)
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

func ppanic(logger *log.Logger, input LogInput) {
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
	var msgs []string
	var errs []string

	for _, arg := range args {
		if argFields, ok := arg.(map[string]interface{}); ok {
			for k, v := range argFields {
				fields[k] = v
			}
		}

		if argString, ok := arg.(string); ok {
			msgs = append(msgs, argString)
		}

		if _, ok := fields["stack"]; ok {
			continue
		}

		if argErr, ok := arg.(error); ok {
			errs = append(errs, argErr.Error())
			fields["stacktrace"] = fmt.Sprintf("%+v", errors.WithStack(argErr))
		}
	}

	var msgStr = strings.Join(msgs, "; ")
	var errStr = strings.Join(errs, "; ")

	if msgStr != "" && errStr != "" {
		return LogInput{Msg: fmt.Sprintf("%v: %v", msgStr, errStr), Params: fields}
	} else if errStr != "" {
		return LogInput{Msg: errStr, Params: fields}
	} else if msgStr != "" {
		return LogInput{Msg: msgStr, Params: fields}
	}

	return LogInput{Msg: msgStr, Params: fields}
}
