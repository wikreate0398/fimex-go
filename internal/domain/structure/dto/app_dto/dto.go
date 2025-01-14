package app_dto

import (
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
	domain_services "wikreate/fimex/internal/domain/services"
	"wikreate/fimex/internal/infrastructure/database/repositories"
)

type AppDeps struct {
	Repository *repositories.Repositories
	Services   *domain_services.Services
	//AppServicess
	Config *config.Config
	Logger interfaces.Logger
}

type Application struct {
	Deps *AppDeps
}
