package logger

import (
	"log/slog"
	"os"
)

// InitLogger initializes the global logger with sane defaults.
// It uses a JSON handler for structured logging, which is ideal for production.
// The log level can be dynamically set via the APP_LEVEL environment variable.
//
// Supported Levels: DEBUG, INFO, WARN, ERROR
func InitLogger() {
	var level slog.Level
	switch os.Getenv("APP_LEVEL") {
	case "DEBUG":
		level = slog.LevelDebug
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
		// AddSource: true, // Uncomment to include file and line number in logs
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)

	slog.Info("Logger initialized", "level", level.String())
}
