package day2

import (
	"fmt"
	"os"
)

// Part1 ...
func Part1() {
	file, err := os.Open("data/day2/full.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	games := parseFile(file)
	fmt.Printf("Num. games parsed: %d\n", len(games))
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

	fmt.Printf("Possible games: %v\n", possibleGames)
	sum := 0
	for _, g := range possibleGames {
		sum += g + 1
	}
	fmt.Println(sum)
}
