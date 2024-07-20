package config

import (
	"log/slog"
	"os"
)

func logFatalError(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}

func logInfo(msg string, args ...any) {
	slog.Info(msg, args...)
}
