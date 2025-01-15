package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var (
	logsDir = "logs"
)

func NewLogger() (*LoggerManager, error) {
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
	logger.SetLevel(log.TraceLevel)
	logger.SetOutput(io.MultiWriter(file, os.Stdout))

	return &LoggerManager{logger: logger}, nil
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
