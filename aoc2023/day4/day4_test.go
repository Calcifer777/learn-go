package day4

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1Sample(t *testing.T) {
	// t.Skip("Skip for now")
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	out, e := Part1("testdata/part1.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 13
	assert.Equal(t, expected, out, "should match")
}

// func TestDay3Part1Custom(t *testing.T) {
// 	t.Skip("Skip for now")
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
// 	slog.SetDefault(logger)
// 	_, e := Part1("testdata/part1-custom.txt")
// 	assert.Nil(t, e, "Part1 failed!")
// }

func TestDay3Part1Full(t *testing.T) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	out, e := Part1("testdata/full.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 27059
	assert.Equal(t, expected, out, "should match")
}

// func TestDay3Part2Sample(t *testing.T) {
// 	t.Skip("Skip for now")
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
// 	slog.SetDefault(logger)
// 	out, e := Part2("testdata/part1.txt")
// 	assert.Nil(t, e, "Part1 failed!")
// 	expected := 467835
// 	assert.Equal(t, expected, out, "should match")
// }
//
// func TestDay3Part2Full(t *testing.T) {
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelWarn,
// 	}
// 	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
// 	slog.SetDefault(logger)
// 	out, e := Part2("testdata/full.txt")
// 	assert.Nil(t, e, "Part1 failed!")
// 	expected := 75805607
// 	assert.Equal(t, expected, out, "should match")
// }
//
