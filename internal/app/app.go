package app

import (
	"context"
	"wikreate/fimex/internal/config"
	domain_serivces "wikreate/fimex/internal/domain/services"
	"wikreate/fimex/internal/dto/app_dto"
	"wikreate/fimex/internal/infrastructure/database/repositories"
	"wikreate/fimex/internal/transport/messagebus"
	"wikreate/fimex/internal/transport/rest"
	"wikreate/fimex/pkg/database"
	"wikreate/fimex/pkg/database/mysql"
	"wikreate/fimex/pkg/lifecycle"
	"wikreate/fimex/pkg/logger"
)

func NewApplication(deps *app_dto.AppDeps) *app_dto.Application {
	return &app_dto.Application{Deps: deps}
}

func Make(cfg *config.Config, log *logger.LoggerManager) {
	ctx := context.Background()

	dbConf := cfg.Databases.MySql
	dbManager := database.NewDBManager(ctx, mysql.DBCreds{
		Host:     dbConf.Host,
		Port:     dbConf.Port,
		User:     dbConf.User,
		Password: dbConf.Password,
		Database: dbConf.Database,
	}, log)

	repo := repositories.NewRepositories(dbManager)
	serv := domain_serivces.NewServices(repo)

	app := NewApplication(&app_dto.AppDeps{
		Repository: repo,
		Services:   serv,
		Config:     cfg,
		Logger:     log,
	})

	lf := lifecycle.Register(
		rest.Init(app),
		messagebus.Init(app),
	)

	lf.Run()
}
