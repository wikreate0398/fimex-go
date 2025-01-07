package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wikreate/fimex/internal/domain/core"
)

type Json map[string]interface{}

type BaseController struct {
	application *core.Application
}

func (ctrl *BaseController) ok200(context *gin.Context, result map[string]any) {
	context.JSON(http.StatusOK, result)
}

func (ctrl *BaseController) error400(context *gin.Context, result map[string]any) {
	context.JSON(http.StatusBadRequest, result)
}
