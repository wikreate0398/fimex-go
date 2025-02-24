package config

import "go.uber.org/fx"

var Provider = fx.Provide(NewConfig)
