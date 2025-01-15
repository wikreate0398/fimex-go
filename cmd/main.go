package main

import (
	"fmt"
	"os"
	"path/filepath"
	"wikreate/fimex/internal/app"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/pkg/logger"
)

func getLogDir() (string, error) {
	// Получение пути к исполняемому файлу
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	// Директория, где лежит бинарник
	execDir := filepath.Dir(execPath)

	// Путь к директории логов относительно бинарника
	logDir := filepath.Join(execDir, "logs")
	return logDir, nil
}

func main() {
	log, err := logger.NewLogger()

	dir, _ := getLogDir()

	fmt.Println(dir)

	if err != nil {
		fmt.Println("Error initializing logger:", err)
		os.Exit(1)
	}

	cfg, err := config.Init("stage")
	fmt.Println(err)
	log.FatalOnErr(err, "Failed to initialize config")

	log.Error("lorems")

	app.Make(cfg, log)
}
