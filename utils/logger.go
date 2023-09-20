package utils

import (
	"cars-go/config"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

func createLogsFolder() {
	newPath := filepath.Join(".", config.LOGS_FOLDER)
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating logs folder: %v", err)
		panic(err)
	}
}

func writeLogsToFile() {
	createLogsFolder()

	file, err := os.OpenFile(config.LOGS_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error at opening logs file: %v", err)
	}
	defer file.Close()

	wrt := io.MultiWriter(os.Stdout, file)
	log.SetOutput(wrt)
}

func NewLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	slog.SetDefault(logger)
	slog.LogAttrs(slog.LevelInfo, "Processed item")
}
