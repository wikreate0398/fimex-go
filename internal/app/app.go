package app

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"strings"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/dto/app_dto"
	"wikreate/fimex/internal/helpers"
	"wikreate/fimex/pkg/database"
	"wikreate/fimex/pkg/database/mysql"
	"wikreate/fimex/pkg/logger"
)

func NewApplication(deps app_dto.AppDeps) *app_dto.Application {
	return &app_dto.Application{Deps: deps}
}

func Test(logger2 interfaces.Logger) error {
	fmt.Println("Test")
	return errors.New("Test")
}

type LogrusLogger struct {
	Logger interfaces.Logger
}

// LogEvent handles a log event for fx application container
func (l *LogrusLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Debug("on start hook executing", helpers.KeyValue{
			"callee": e.FunctionName,
			"caller": e.CallerName,
		})
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Errorf("on start hook failed: %v", e.Err, helpers.KeyValue{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			})
		} else {
			l.Logger.Debug("on start hook executed", helpers.KeyValue{
				"callee":  e.FunctionName,
				"caller":  e.CallerName,
				"runtime": e.Runtime.String(),
			})
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Debug("on stop hook executing", helpers.KeyValue{
			"callee": e.FunctionName,
			"caller": e.CallerName,
		})
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Errorf("on stop hook failed: %v", e.Err, helpers.KeyValue{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			})
		} else {
			l.Logger.Debug("on stop hook executed", helpers.KeyValue{
				"callee":  e.FunctionName,
				"caller":  e.CallerName,
				"runtime": e.Runtime.String(),
			})
		}
	case *fxevent.Supplied:
		l.Logger.Debugf("supplied: %v", e.Err, helpers.KeyValue{
			"type":   e.TypeName,
			"module": e.ModuleName,
		})
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("provided", helpers.KeyValue{
				"constructor": e.ConstructorName,
				"module":      e.ModuleName,
				"type":        rtype,
			})
		}
		if e.Err != nil {
			l.Logger.Errorf("error encountered while applying options: %v", e.Err, helpers.KeyValue{
				"module": e.ModuleName,
			})
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("replaced", helpers.KeyValue{
				"module": e.ModuleName,
				"type":   rtype,
			})
		}
		if e.Err != nil {
			l.Logger.Errorf("error encountered while replacing: %v", e.Err, helpers.KeyValue{
				"module": e.ModuleName,
			})
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("decorated", helpers.KeyValue{
				"module": e.ModuleName,
				"type":   rtype,
			})
		}
		if e.Err != nil {
			l.Logger.Errorf("error encountered while applying options: %v", e.Err, helpers.KeyValue{
				"module": e.ModuleName,
			})
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.
		l.Logger.Debug("invoking", helpers.KeyValue{
			"function": e.FunctionName,
			"module":   e.ModuleName,
		})
	case *fxevent.Invoked:
		if e.Err != nil {
			l.Logger.WithFields(helpers.KeyStrValue{
				"stack":    e.Trace,
				"function": e.FunctionName,
				"module":   e.ModuleName,
			}).Errorf("invoke failed: %v", e.Err)
		}
	case *fxevent.Stopping:
		l.Logger.Debugf("received signal: %s", strings.ToUpper(e.Signal.String()))
	case *fxevent.Stopped:
		if e.Err != nil {
			l.Logger.Errorf("received signal: %v", e.Err)
		}
	case *fxevent.RollingBack:
		l.Logger.Errorf("start failed, rolling back: %v", e.StartErr)
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.Logger.Errorf("rollback failed: %v", e.Err)
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.Logger.Errorf("start failed: %v", e.Err)
		} else {
			l.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Logger.Errorf("custom logger initialization failed: %v", e.Err)
		} else {
			l.Logger.Debug("initialized custom fxevent.Logger", helpers.KeyValue{
				"function": e.ConstructorName,
			})
		}
	}
}

func Make() {
	fx.New(
		config.Module,
		logger.Module,

		fx.Invoke(Test),
		fx.Invoke(func(lc fx.Lifecycle) {
			fmt.Println("hello")
		}),

		fx.WithLogger(func(log interfaces.Logger) fxevent.Logger {
			return &LogrusLogger{
				Logger: log,
			}
		}),
	).Run()

	return
	//ctx := context.Background()
	//
	//dbConf := cfg.Databases.MySql
	//dbManager := database.NewMysqlManager(ctx, mysql.DBCreds{
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

func ProvideMysql(ctx context.Context, cfg *config.Config, log interfaces.Logger) *database.DB {
	dbConf := cfg.Databases.MySql
	db, err := database.NewMysqlManager(ctx, mysql.DBCreds{
		Host:     dbConf.Host,
		Port:     dbConf.Port,
		User:     dbConf.User,
		Password: dbConf.Password,
		Database: dbConf.Database,
	})

	log.FatalOnErr(err, "Failed to connect to database")

	return db
}

/**
config
logging
database

repository
domain services

http
messagebuss

*/
