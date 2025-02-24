package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/transport/rest/controllers"
)

type RoutesParams struct {
	fx.In

	Logger interfaces.Logger
	Router *gin.Engine

	MainController *controllers.MainController
}

func newRouter() *gin.Engine {
	return gin.Default()
}

func registerRoutes(p RoutesParams) {
	v1 := p.Router.Group("/v1")
	{
		v1.GET("/", p.MainController.Home)
	}
}
