package config

import (
	"log/slog"
	"os"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func LoadLogger(env string) *slog.Logger {
	logLevel := slog.LevelInfo

	switch env {
	case envDev:
		logLevel = slog.LevelDebug
		break
	case envProd:
		logLevel = slog.LevelInfo
		break
	}

	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}),
	)
}
