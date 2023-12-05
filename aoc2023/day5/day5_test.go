package day5

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	t.Skip("skip")
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	out, e := Part1("testdata/part1.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 35
	assert.NotNil(t, out)
	assert.Equal(t, expected, out, "should match")
}

func TestPart1Full(t *testing.T) {
	t.Skip("skip")
	opts := &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	out, e := Part1("testdata/full.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 309796150
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Sample(t *testing.T) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	out, e := Part2("testdata/part1.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 0
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Full(t *testing.T) {
	t.Skip("Skip for now")
	opts := &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	out, e := Part2("testdata/full.txt")
	assert.Nil(t, e, "Part2 failed!")
	expected := 5744979
	assert.Equal(t, expected, out, "should match")
}
