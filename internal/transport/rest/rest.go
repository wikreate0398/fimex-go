package rest

import (
	"context"
	"go.uber.org/fx"
	"wikreate/fimex/internal/transport/rest/server"
)

type ServerParams struct {
	fx.In

	Server *server.Server
	Lc     fx.Lifecycle
}

func handleServer(p ServerParams) {
	p.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			p.Server.Start()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			p.Server.Stop(ctx)
			return nil
		},
	})
}
