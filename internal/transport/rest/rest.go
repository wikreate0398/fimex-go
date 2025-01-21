package rest

import (
	"context"
	"net/http"
	"wikreate/fimex/internal/dto/app_dto"
	"wikreate/fimex/pkg/lifecycle"
	"wikreate/fimex/pkg/server"
)

func Init(app *app_dto.Application) func(lf *lifecycle.Lifecycle) {
	return func(lf *lifecycle.Lifecycle) {

		obj := server.NewServer(InitRouter(app), app.Deps.Config)

		lf.Append(lifecycle.AppendLifecycle{
			OnStart: func(ctx context.Context) any {
				if err := obj.Start(); err != nil && err != http.ErrServerClosed {
					app.Deps.Logger.Error(err)
				}
				return nil
			},

			OnStop: func(ctx context.Context) any {
				err := obj.Stop(ctx)
				app.Deps.Logger.PanicOnErr(err, "Failed to stop services")
				return nil
			},
		})
	}
}
