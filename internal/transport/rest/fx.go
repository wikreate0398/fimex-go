package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"wikreate/fimex/internal/transport/rest/controllers"
)

var Module = fx.Module("rest",

	fx.Provide(gin.Default),
	controllers.Module,

	fx.Invoke(RegisterRoutes),
	fx.Invoke(BootstrapServer),
)
