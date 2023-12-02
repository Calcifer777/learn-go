package main

import (
	"fmt"
	"log/slog"
)

func main() {
	fmt.Println("Hello, world")
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
}
