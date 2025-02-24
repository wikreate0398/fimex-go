package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
	domain_services "wikreate/fimex/internal/domain/services"
	"wikreate/fimex/internal/infrastructure/db"
	"wikreate/fimex/internal/infrastructure/logger"
	"wikreate/fimex/internal/infrastructure/storage/repositories"
	"wikreate/fimex/internal/transport/rbbtmq"
	"wikreate/fimex/internal/transport/rest"
)

func Create() {
	fx.New(
		config.Provider,
		logger.Provider,
		db.Provider,

		repositories.Module,
		domain_services.Module,

		fx.Options(
			rest.Module,
			rbbtmq.Module,
		),

		fx.WithLogger(func(log interfaces.Logger) fxevent.Logger {
			return logger.NewFxLogger(log)
		}),
	).Run()
}
