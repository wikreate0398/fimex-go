package app

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/dto/app_dto"
	"wikreate/fimex/pkg/database"
	"wikreate/fimex/pkg/database/mysql"
	"wikreate/fimex/pkg/logger"
)

func NewApplication(deps app_dto.AppDeps) *app_dto.Application {
	return &app_dto.Application{Deps: deps}
}

func Test(logger2 interfaces.Logger) error {
	fmt.Println("Test")
	return nil
}

func Make() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			func() interfaces.Logger {
				obj, _ := logger.NewLogger()
				return obj
			},
		),
		fx.Invoke(Test),
		fx.Invoke(func(lc fx.Lifecycle) {
			fmt.Println("hello")
		}),
	).Run()

	return
	//ctx := context.Background()
	//
	//dbConf := cfg.Databases.MySql
	//dbManager := database.NewDBManager(ctx, mysql.DBCreds{
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

func ProvideMysql(ctx context.Context, cfg *config.Config, log interfaces.Logger) *database.DbAdapter {
	dbConf := cfg.Databases.MySql
	return database.NewDBManager(ctx, mysql.DBCreds{
		Host:     dbConf.Host,
		Port:     dbConf.Port,
		User:     dbConf.User,
		Password: dbConf.Password,
		Database: dbConf.Database,
	}, log)
}
