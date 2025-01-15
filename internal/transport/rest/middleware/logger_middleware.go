package middleware

import (
	"github.com/gin-gonic/gin"
	"wikreate/fimex/internal/domain/structure/dto/app_dto"
	"wikreate/fimex/internal/helpers"
)

func LoggerMiddleware(app *app_dto.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Логируем ошибки, если они есть
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				app.Deps.Logger.Error(err, "Logger Middleware", helpers.KeyValue{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				})
			}
		}
	}
}
