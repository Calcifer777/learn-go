package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calcifer777/aoc2023/day1"
	"github.com/calcifer777/aoc2023/day2"
	flag "github.com/spf13/pflag"
)

type Selection struct {
	day, part int
}

func NewSelection(day, part int) Selection {
	return Selection{day, part}
}

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	day := flag.IntP("day", "d", 1, "Day to run")
	part := flag.IntP("part", "p", 1, "Part to run [1|2]")
	selection := NewSelection(*day, *part)
	flag.Parse()
	var out int
	switch selection {
	case Selection{1, 1}:
		out = day1.Part1("day1/testdata/full.txt")
	case Selection{1, 2}:
		out = day1.Part2("day1/testdata/full.txt")
	case Selection{2, 1}:
		out = day2.Part1("day2/testdata/full.txt")
	case Selection{2, 2}:
		out = day2.Part2("day2/testdata/full.txt")
	default:
		fmt.Printf("Error: Day %d\n", *day)
	}

	slog.Warn("Success", slog.Int("day", *day), slog.Int("part", *part), slog.Int("Result", out))

}
