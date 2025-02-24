package rest

import (
	"go.uber.org/fx"
	"wikreate/fimex/internal/transport/rest/controllers"
	"wikreate/fimex/internal/transport/rest/server"
)

var Module = fx.Module("rest",
	fx.Provide(
		fx.Private,

		newRouter,
		server.NewServer,

		controllers.NewMainController,
	),

	fx.Invoke(
		registerMiddleware,
		registerRoutes,
		handleServer,
	),
)
