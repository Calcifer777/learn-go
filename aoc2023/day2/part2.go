package day2

import (
	"fmt"
	"os"
)

// Part2 ...
func Part2() {
	file, err := os.Open("data/day2/full.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	games := parseFile(file)
	fmt.Printf("Num. games parsed: %d\n", len(games))
	bags := make([]Bag, 0)
	for _, g := range games {
		bag := Bag{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, draw := range g {
			for color, n := range draw {
				bagN, ok := bag[color]
				if !ok {
					panic("Somehing went wrong")
				}
				if bagN < n {
					bag[color] = n
				}
			}
		}
		bags = append(bags, bag)
	}

	sum := 0
	for idx, b := range bags {
		power := b["red"] * b["green"] * b["blue"]
		sum += power
		fmt.Printf("%d: power %d\n", idx, power)
	}
	fmt.Printf("Sum: %d\n", sum)
}
