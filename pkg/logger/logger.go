package logger

import (
	"log/slog"
	"os"
)


func New() *slog.Logger {
	opc := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(os.Stdout, opc)
	logger := slog.New(handler)
	
	return logger
}