package messagebus

import (
	"context"
	"fmt"
	"sync"
	"wikreate/fimex/internal/domain/core"
	"wikreate/fimex/internal/transport/messagebus/consumers"
	"wikreate/fimex/pkg/lifecycle"
	"wikreate/fimex/pkg/rabbitmq"
)

func Init(application *core.Application) func(lf *lifecycle.Lifecycle) {
	return func(lf *lifecycle.Lifecycle) {

		lf.Append(lifecycle.AppendLifecycle{
			OnStart: func(ctx context.Context) any {
				fmt.Println("Message bus Init")

				ctx, cancel := context.WithCancel(ctx)
				queues := consumers.NewHandlers(application)

				conf := application.Config.RabbitMQ
				rbMq := rabbitmq.InitRabbitMQ(rabbitmq.Credentials{
					Host:     conf.Host,
					Port:     conf.Port,
					User:     conf.User,
					Password: conf.Password,
				})

				defer rbMq.Close()

				wg := new(sync.WaitGroup)

				// Listners...
				rbMq.Listen(rabbitmq.ListnerInput{
					Ctx:        ctx,
					Exchange:   "catalog",
					QueueName:  "generate_names",
					RoutingKey: "",
					Resolver:   queues.GenerateNamesQueue,
					Wg:         wg,
				})
				
				select {
				case <-ctx.Done():
					cancel()
					fmt.Println("Message bus has been stopped")
				}

				wg.Wait()

				return nil
			},

			OnStop: func(ctx context.Context) any {
				return nil
			},
		})
	}
}
