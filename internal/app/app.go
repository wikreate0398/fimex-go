package app

import (
	"context"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/core"
	"wikreate/fimex/internal/repository"
	"wikreate/fimex/internal/services"
	"wikreate/fimex/internal/transport/messagebus"
	"wikreate/fimex/internal/transport/rest"
	"wikreate/fimex/pkg/database"
	"wikreate/fimex/pkg/database/mysql"
	"wikreate/fimex/pkg/lifecycle"
)

func NewApplication(deps core.AppDeps) *core.Application {
	return &core.Application{AppDeps: deps}
}

func Make(cfg *config.Config) {

	ctx := context.Background()

	dbConf := cfg.Databases.MySql
	dbManager := database.NewDBManager(ctx, mysql.DBCreds{
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Database,
	})

	repo := repository.NewRepository(dbManager)
	serv := services.NewService(repo)

	app := NewApplication(core.AppDeps{
		Repository: repo,
		Service:    serv,
		Config:     cfg,
	})

	lf := lifecycle.Register(
		rest.Init(app),
		messagebus.Init(app),
	)

	lf.Run()
}
