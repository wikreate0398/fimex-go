package main

import (
	"fmt"
	"os"
	"wikreate/fimex/internal/app"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/pkg/logger"
)

func main() {
	log, err := logger.NewLogger()

	if err != nil {
		fmt.Println("Error initializing logger:", err)
		os.Exit(1)
	}

	cfg, err := config.Init("stage")
	log.FatalOnErr(err, "Failed to initialize config")

	app.Make(cfg, log)
}
