package main

import (
	"log/slog"
	"os"
)

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}
