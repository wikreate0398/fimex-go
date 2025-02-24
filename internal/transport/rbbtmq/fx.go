package rbbtmq

import (
	"go.uber.org/fx"
	"wikreate/fimex/internal/transport/rbbtmq/consumers"
)

var Module = fx.Module("rbbtmq",
	fx.Provide(
		fx.Private,

		consumers.NewGenerateProductsNamesConsumer,
		consumers.NewRecalcHistoryBallanceConsumer,
		consumers.NewSortProductsConsumer,
	),

	fx.Invoke(handleRabbitMq),
)
