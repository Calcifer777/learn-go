package day3

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1Sample(t *testing.T) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	out, e := Part1("testdata/part1.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 4361
	assert.Equal(t, expected, out, "should match")
}

func TestDay3Part1Full(t *testing.T) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	// t.Skip("Skip for now")
	out, e := Part1("testdata/full.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := -1
	assert.Equal(t, expected, out, "should match")
}
