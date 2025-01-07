package rest

import (
	"context"
	"net/http"
	"wikreate/fimex/internal/domain/core"
	"wikreate/fimex/pkg/failed"
	"wikreate/fimex/pkg/lifecycle"
	"wikreate/fimex/pkg/logger"
	"wikreate/fimex/pkg/server"
)

func Init(app *core.Application) func(lf *lifecycle.Lifecycle) {
	return func(lf *lifecycle.Lifecycle) {

		obj := server.NewServer(InitRouter(app), app.Config)

		lf.Append(lifecycle.AppendLifecycle{
			OnStart: func(ctx context.Context) any {
				if err := obj.Start(); err != nil && err != http.ErrServerClosed {
					logger.Error(logger.LogInput{Msg: err})
				}
				return nil
			},

			OnStop: func(ctx context.Context) any {
				err := obj.Stop(ctx)
				failed.PanicOnError(err, "Failed to stop services")
				return nil
			},
		})
	}
}
