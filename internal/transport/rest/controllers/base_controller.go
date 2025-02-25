package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Json map[string]interface{}

type BaseController struct {
}

func (ctrl *BaseController) ok200(context *gin.Context, result map[string]any) {
	context.JSON(http.StatusOK, result)
}

// nolint:unused
func (ctrl *BaseController) error400(context *gin.Context, result map[string]any) {
	context.JSON(http.StatusBadRequest, result)
}
