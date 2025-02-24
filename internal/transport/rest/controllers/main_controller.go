package controllers

import (
	"github.com/gin-gonic/gin"
)

type MainController struct {
	BaseController
}

func NewMainController() *MainController {
	return &MainController{
		BaseController{},
	}
}

func (c *MainController) Home(ctx *gin.Context) {
	c.ok200(ctx, Json{"data": "lorem"})
}
