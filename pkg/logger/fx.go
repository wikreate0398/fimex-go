package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"wikreate/fimex/internal/domain/interfaces"
)

var Module = fx.Options(
	fx.Provide(func() interfaces.Logger {
		logManger, err := NewLogger()

		if err != nil {
			log.Fatal(fmt.Sprintf("Error initializing logger: %v", err))
		}

		return logManger
	}),
)
