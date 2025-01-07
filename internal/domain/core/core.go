package core

import (
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/repository"
	"wikreate/fimex/internal/services"
)

type AppDeps struct {
	Repository *repository.Repository
	Service    *services.Service
	Config     *config.Config
}

type Application struct {
	AppDeps
}
