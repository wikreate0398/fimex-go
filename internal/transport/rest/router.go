package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/transport/rest/controllers"
	"wikreate/fimex/internal/transport/rest/middleware"
)

type Params struct {
	fx.In

	Logger interfaces.Logger
	Router *gin.Engine

	MainController *controllers.MainController
}

func RegisterRoutes(p Params) {

	p.Router.Use(middleware.LoggerMiddleware(p.Logger))
	p.Router.Use(gin.Recovery())

	v1 := p.Router.Group("/v1")
	{
		v1.GET("/", p.MainController.Home)
	}
}
