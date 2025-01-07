package controllers

import (
	"wikreate/fimex/internal/domain/core"
)

type Controllers struct {
	MainController *MainController
}

func NewControllers(application *core.Application) *Controllers {
	return &Controllers{
		MainController: NewMainController(application),
	}
}
