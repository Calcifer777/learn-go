package main

import (
	"fmt"
	"strconv"
	"utils"
)

func Part1() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	var prev int = -1
	var i int
	for _, line := range lines {
		curr, _ := strconv.Atoi(line)
		if prev != -1 && curr > prev {
			i++
		}
		prev = curr
	}
	fmt.Printf("Part 1: %d\n", i)
}

func Part2() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	var prev int
	var i int
	for idx, line := range lines {
		if idx < 3 {
			continue
		}
		curr, _ := strconv.Atoi(line)
		prev, _ = strconv.Atoi(lines[idx-3])
		if curr > prev {
			i++
		}
		prev = curr
	}
	fmt.Printf("Part 2: %d\n", i)
}

func main() {
	Part2()
}
