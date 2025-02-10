package messagebus

import (
	"context"
	"sync"
	"wikreate/fimex/internal/dto/app_dto"
	"wikreate/fimex/internal/transport/messagebus/consumers"
	"wikreate/fimex/pkg/lifecycle"
	"wikreate/fimex/pkg/rabbitmq"
)

func Init(application *app_dto.Application) func(lf *lifecycle.Lifecycle) {
	return func(lf *lifecycle.Lifecycle) {

		lf.Append(lifecycle.AppendLifecycle{
			OnStart: func(ctx context.Context) any {
				application.Deps.Logger.Info("Message bus Init")

				queues := consumers.NewHandlers(application)

				conf := application.Deps.Config.RabbitMQ
				rbMq := rabbitmq.InitRabbitMQ(rabbitmq.Credentials{
					Host:     conf.Host,
					Port:     conf.Port,
					User:     conf.User,
					Password: conf.Password,
				}, application.Deps.Logger)

				defer rbMq.Close()

				wg := new(sync.WaitGroup)

				rbMq.Register(rabbitmq.RegisterDto{
					Exchange:   "catalog",
					QueueName:  "products_queue",
					RoutingKey: "generate.names",
					Resolver:   queues.GenerateProductsNamesQueue,
				})

				rbMq.Register(rabbitmq.RegisterDto{
					Exchange:   "catalog",
					QueueName:  "products_queue",
					RoutingKey: "sort.product",
					Resolver:   queues.SortProductsQueue,
				})

				rbMq.Register(rabbitmq.RegisterDto{
					Exchange:   "payment",
					QueueName:  "ballance_queue",
					RoutingKey: "recalculate.history",
					Resolver:   queues.RecalcBallanceHistory,
				})

				rbMq.Listen(ctx, wg)

				<-ctx.Done()
				wg.Wait()

				return nil
			},

			OnStop: func(ctx context.Context) any {
				return nil
			},
		})
	}
}
