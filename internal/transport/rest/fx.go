package rest

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module("http", fx.Invoke(Init))
}
