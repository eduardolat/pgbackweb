package logger

import (
	"log/slog"
	"os"
	"sync"
)

var logger *slog.Logger
var loggerOnce sync.Once

// getLogger returns a logger singleton configured to log to
// custom writer in JSON format
func getLogger() *slog.Logger {
	loggerOnce.Do(func() {
		w := &logWriter{}
		logger = slog.New(slog.NewJSONHandler(
			w, nil,
		))
	})

	return logger
}

// Debug logs a debug message
func Debug(msg string, args ...KV) {
	getLogger().Debug(msg, kvToArgs(args...)...)
}

// DebugNs logs a debug message with a namespace
// It's for grouping logs consistently
func DebugNs(ns string, msg string, args ...KV) {
	getLogger().Debug(msg, kvToArgsNs(ns, args...)...)
}

// Info logs a info message
func Info(msg string, args ...KV) {
	getLogger().Info(msg, kvToArgs(args...)...)
}

// InfoNs logs a info message with a namespace
// It's for grouping logs consistently
func InfoNs(ns string, msg string, args ...KV) {
	getLogger().Info(msg, kvToArgsNs(ns, args...)...)
}

// Warn logs a warn message
func Warn(msg string, args ...KV) {
	getLogger().Warn(msg, kvToArgs(args...)...)
}

// WarnNs logs a warn message with a namespace
// It's for grouping logs consistently
func WarnNs(ns string, msg string, args ...KV) {
	getLogger().Warn(msg, kvToArgsNs(ns, args...)...)
}

// Error logs a error message
func Error(msg string, args ...KV) {
	getLogger().Error(msg, kvToArgs(args...)...)
}

// ErrorNs logs a error message with a namespace
// It's for grouping logs consistently
func ErrorNs(ns string, msg string, args ...KV) {
	getLogger().Error(msg, kvToArgsNs(ns, args...)...)
}

// FatalError is equivalent to Error() followed by a call to os.Exit(1)
func FatalError(msg string, args ...KV) {
	Error(msg, args...)
	os.Exit(1)
}

// FatalErrorNs is equivalent to FatalError() with a namespace
func FatalErrorNs(ns string, msg string, args ...KV) {
	ErrorNs(ns, msg, args...)
	os.Exit(1)
}
