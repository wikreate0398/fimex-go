package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"wikreate/fimex/internal/app"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/pkg/logger"
)

func main() {
	logManger, err := logger.NewLogger()

	if err != nil {
		log.Fatal(fmt.Sprintf("Error initializing logger: %v", err))
	}

	cfg, err := config.Init("stage")
	logManger.FatalOnErr(err, "Failed to initialize config")

	app.Make(cfg, logManger)
}
