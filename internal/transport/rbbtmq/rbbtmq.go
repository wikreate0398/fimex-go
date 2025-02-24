package rbbtmq

import (
	"context"
	"go.uber.org/fx"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/transport/rbbtmq/consumers"
	"wikreate/fimex/pkg/rabbitmq"
)

type RabitMqParams struct {
	fx.In

	Lc     fx.Lifecycle
	Logger interfaces.Logger
	Config *config.Config

	GenerateProductsNamesHandler *consumers.GenerateNamesConsumer
	SortProductsHandler          *consumers.SortProductsConsumer
	RecalcBallanceHistoryHandler *consumers.RecalcHistoryBallanceConsumer
}

func handleRabbitMq(p RabitMqParams) {
	conf := p.Config.RabbitMQ
	rbMq := rabbitmq.InitRabbitMQ(rabbitmq.Credentials{
		Host:     conf.Host,
		Port:     conf.Port,
		User:     conf.User,
		Password: conf.Password,
	}, p.Logger)

	p.Lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {

			rbMq.Register(rabbitmq.RegisterDto{
				Exchange:   "catalog",
				QueueName:  "products_queue",
				RoutingKey: "generate.names",
				Resolver:   p.GenerateProductsNamesHandler,
			})

			rbMq.Register(rabbitmq.RegisterDto{
				Exchange:   "catalog",
				QueueName:  "products_queue",
				RoutingKey: "sort.product",
				Resolver:   p.SortProductsHandler,
			})

			rbMq.Register(rabbitmq.RegisterDto{
				Exchange:   "payment",
				QueueName:  "ballance_queue",
				RoutingKey: "recalculate.history",
				Resolver:   p.RecalcBallanceHistoryHandler,
			})

			rbMq.Listen()

			return nil
		},

		OnStop: func(_ context.Context) error {
			rbMq.Close()
			return nil
		},
	})
}
