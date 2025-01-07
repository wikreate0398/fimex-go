package controllers

import (
	"github.com/gin-gonic/gin"
	"wikreate/fimex/internal/domain/core"
)

type MainController struct {
	BaseController
}

func NewMainController(application *core.Application) *MainController {
	return &MainController{
		BaseController{application},
	}
}

func (c *MainController) Home(ctx *gin.Context) {
	c.ok200(ctx, Json{"data": "lorem"})
}
