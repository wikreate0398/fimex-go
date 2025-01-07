package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	fileName := fmt.Sprintf("logs/logs-%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Warn("Не удалось открыть файл логов, используется вывод в консоль")
	}
}

type LogFields log.Fields

type LogInput struct {
	Msg    interface{}
	Params LogFields
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
