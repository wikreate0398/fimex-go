package controllers

import (
	"wikreate/fimex/internal/dto/app_dto"
)

type Controllers struct {
	MainController *MainController
}

func NewControllers(application *app_dto.Application) *Controllers {
	return &Controllers{
		MainController: NewMainController(application),
	}
}
