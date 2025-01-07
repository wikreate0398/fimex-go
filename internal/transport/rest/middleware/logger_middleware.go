package middleware

import (
	"github.com/gin-gonic/gin"
	"wikreate/fimex/pkg/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Логируем ошибки, если они есть
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.Error(logger.LogInput{Msg: err.Error(), Params: logger.LogFields{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}})
			}
		}
	}
}
