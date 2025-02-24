package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/infrastructure/db"
	"wikreate/fimex/internal/infrastructure/logger"
	"wikreate/fimex/internal/infrastructure/storage/repositories"
	"wikreate/fimex/internal/transport/rest"
)

//func NewApplication(deps app_dto.AppDeps) *app_dto.Application {
//	return &app_dto.Application{Deps: deps}
//}

func Test(logger interfaces.Logger, db interfaces.DB) error {
	user := struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}{}
	db.Get(&user, "SELECT id, name FROM users LIMIT 1")

	return nil
}

func Create() {
	fx.New(
		config.Provider,
		logger.Provider,
		db.Provider,

		repositories.Module,

		rest.Module,

		fx.WithLogger(func(log interfaces.Logger) fxevent.Logger {
			return logger.NewFxLogger(log)
		}),
	).Run()

	return
	//ctx := context.Background()
	//
	//dbConf := cfg.Databases.MySql
	//dbManager := storage.NewDb(ctx, mysql.DBCreds{
	//	Host:     dbConf.Host,
	//	Port:     dbConf.Port,
	//	User:     dbConf.User,
	//	Password: dbConf.Password,
	//	Database: dbConf.Database,
	//}, log)
	//
	//repo := repositories.NewRepositories(dbManager)
	//serv := domain_serivces.NewServices(repo, log)
	//
	//app := NewApplication(app_dto.AppDeps{
	//	Repository: repo,
	//	Services:   serv,
	//	Config:     cfg,
	//	Logger:     log,
	//})
	//
	//lf := lifecycle.Register(
	//	rest.Init(app),
	//	messagebus.Init(app),
	//)
	//
	//lf.Run()
}

/**
config
logging
storage

repository
domain services

http
messagebuss

*/
