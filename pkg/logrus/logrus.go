package logrus

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime/debug"
	"time"
)

var (
	logsDir = "logs"
)

type StackHook struct{}

func (h *StackHook) Levels() []log.Level {
	return []log.Level{log.ErrorLevel, log.FatalLevel, log.PanicLevel, log.WarnLevel}
}

func (h *StackHook) Fire(entry *log.Entry) error {
	entry.Data["stack"] = string(debug.Stack())
	return nil
}

func NewLogrus() (*log.Logger, error) {
	if err := createDir(); err != nil {
		return nil, err
	}

	file, err := openFile()

	if err != nil {
		return nil, err
	}

	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetReportCaller(true)
	logger.SetLevel(log.DebugLevel)
	logger.SetOutput(io.MultiWriter(file, os.Stdout))

	logger.AddHook(&StackHook{})

	return logger, nil
}

func createDir() error {
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		if err := os.Mkdir(logsDir, 0755); err != nil {
			return fmt.Errorf("cannot create logs directory: %w", err)
		}

		file, err := os.Create(fmt.Sprintf("%s/.gitignore", logsDir))
		if err != nil {
			return fmt.Errorf("cannot create .gitignore: %w", err)
		}

		defer file.Close()

		if _, err = file.WriteString("!.gitignore\n*"); err != nil {
			return fmt.Errorf("cannot create .gitignore: %w", err)
		}
	}

	return nil
}

func openFile() (*os.File, error) {
	fileName := fmt.Sprintf("%s/logs-%s.log", logsDir, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return file, nil
}
