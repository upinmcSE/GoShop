package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/upinmcSE/goshop/pkg/logger"
)


func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func NewLoggerWithPath(fileName string, level string) *zerolog.Logger {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get working dir:", err)
	}

	path := filepath.Join(cwd, "internal/logs", fileName)

	config := logger.LoggerConfig{
		Level:      level,
		Filename:   path,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     5,
		Compress:   true,
		IsDev:      GetEnv("APP_EVN", "development"),
	}
	return logger.NewLogger(config)
}