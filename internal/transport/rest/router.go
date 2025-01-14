package rest

import (
	"github.com/gin-gonic/gin"
	"wikreate/fimex/internal/domain/structure/dto/app_dto"
	"wikreate/fimex/internal/transport/rest/controllers"
	"wikreate/fimex/internal/transport/rest/middleware"
)

func InitRouter(app *app_dto.Application) *gin.Engine {

	handlers := controllers.NewControllers(app)

	router := gin.New()
	router.Use(middleware.LoggerMiddleware())
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		v1.GET("/", handlers.MainController.Home)
	}

	return router
}
