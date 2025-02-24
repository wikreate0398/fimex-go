package db

import (
	"go.uber.org/fx"
	"wikreate/fimex/internal/domain/interfaces"
)

var Provider = fx.Provide(
	fx.Annotate(
		NewDb,
		fx.As(new(interfaces.DB)),
	),
)
