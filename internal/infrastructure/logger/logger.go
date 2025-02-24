package logger

import (
	"wikreate/fimex/internal/infrastructure/adapters/logger_adapter"
	"wikreate/fimex/pkg/logrus"
)

func NewLogger() (*logger_adapter.LoggerAdapter, error) {
	logger, err := logrus.NewLogrus()

	if err != nil {
		return nil, err
	}

	return logger_adapter.NewLoggerAdapter(logger), nil
}
