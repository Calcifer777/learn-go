package day2

import (
	"fmt"
	"log/slog"
	"os"
)

// Part1 ...
func Part1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	games := parseFile(file)
	slog.Debug(fmt.Sprintf("Num. games parsed: %d\n", len(games)))
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	possibleGames := make([]int, 0)
	for idx, g := range games {
		allowed := true
	drawLoop:
		for _, draw := range g {
			for color, n := range draw {
				bagN, ok := bag[color]
				if (!ok) || (bagN < n) {
					allowed = false
					break drawLoop
				}
			}
		}
		if allowed {
			possibleGames = append(possibleGames, idx)
		}
	}

	slog.Debug(fmt.Sprintf("Possible games: %v\n", possibleGames))
	sum := 0
	for _, g := range possibleGames {
		sum += g + 1
	}
	slog.Debug("Output", slog.Int("sum", sum))
	return sum
}
