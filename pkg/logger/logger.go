package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/AC-Pcong/goscaff/pkg/config"
)

// NewLogger creates a new slog.Logger instance based on the provided LogConfig.
// It is intended to be used as a wire provider.
func NewLogger(cfg *config.Config) *slog.Logger {
	var level slog.Level
	switch cfg.Log.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo // Default to info level
	}

	// For simplicity, we'll use a JSON handler writing to os.Stdout.
	// In a real application, you might want to configure different outputs
	// (e.g., file, syslog) or different formats (e.g., text).
	handlerOptions := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}

	var w io.Writer = os.Stdout
	handler := slog.NewJSONHandler(w, handlerOptions)

	return slog.New(handler)
}
