package main

import (
	"wikreate/fimex/internal/app"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/pkg/failed"
	"wikreate/fimex/pkg/logger"
)

func main() {
	logger.Init()

	cfg, err := config.Init("stage")
	failed.PanicOnError(err, "Failed to initialize config")

	app.Make(cfg)
}
