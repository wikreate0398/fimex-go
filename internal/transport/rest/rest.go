package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/pkg/lifecycle"
	"wikreate/fimex/pkg/server"
)

func BootstrapServer(conf *config.Config, logger interfaces.Logger, router *gin.Engine) func(lf *lifecycle.Lifecycle) {
	return func(lf *lifecycle.Lifecycle) {

		obj := server.NewServer(router, conf)

		lf.Append(lifecycle.AppendLifecycle{
			OnStart: func(ctx context.Context) any {
				if err := obj.Start(); err != nil && err != http.ErrServerClosed {
					logger.Errorf("Failed to start server %v", err)
				}
				return nil
			},

			OnStop: func(ctx context.Context) any {
				err := obj.Stop(ctx)
				logger.Errorf("Failed to stop server %v", err)
				return nil
			},
		})
	}
}
