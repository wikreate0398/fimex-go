package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/helpers"
)

var Provider = fx.Provide(
	fx.Annotate(
		NewLogger,
		fx.As(new(interfaces.Logger)),
	),
)

type FxLogger struct {
	Logger interfaces.Logger
}

func NewFxLogger(log interfaces.Logger) *FxLogger {
	return &FxLogger{Logger: log}
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.WithFields(helpers.KeyStrValue{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			}).Errorf("on start hook failed: %v", e.Err)
		}
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.WithFields(helpers.KeyStrValue{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			}).Errorf("on stop hook failed: %v", e.Err)
		}

	case *fxevent.Provided:
		if e.Err != nil {
			l.Logger.Errorf("error encountered while applying options: %v", e.Err, helpers.KeyValue{
				"module": e.ModuleName,
			})
		}
	case *fxevent.Replaced:
		if e.Err != nil {
			l.Logger.Errorf("error encountered while replacing: %v", e.Err, helpers.KeyValue{
				"module": e.ModuleName,
			})
		}
	case *fxevent.Decorated:
		if e.Err != nil {
			l.Logger.Errorf("error encountered while applying options: %v", e.Err, helpers.KeyValue{
				"module": e.ModuleName,
			})
		}
	case *fxevent.Invoked:
		if e.Err != nil {
			l.Logger.WithFields(helpers.KeyStrValue{
				"function": e.FunctionName,
				"module":   e.ModuleName,
			}).Errorf("invoke failed: %v", e.Err)
		}
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
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Logger.Errorf("custom logrus initialization failed: %v", e.Err)
		}
	}
}
