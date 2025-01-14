package main

import (
	"wikreate/fimex/internal/app"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.Init("stage")
	log.PanicOnFailed(err, "Failed to initialize config")

	app.Make(cfg, log)
}
