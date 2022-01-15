package main

import (
	"fmt"
	"regexp"
	"strconv"
	"utils"
)

func Part1() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	r, err := regexp.Compile(`(?P<Direction>[\w]+)\s(?P<Steps>[\d]+)`)
	var depth, hor int
	for _, line := range lines {
		res := r.FindStringSubmatch(line) // res[0]: direction, res[1]: steps
		var direction = res[1]
		steps, _ := strconv.Atoi(res[2])
		switch direction {
		case "forward":
			hor += steps
		case "backward":
			hor -= steps
		case "up":
			depth -= steps
		case "down":
			depth += steps
		}
	}
	fmt.Printf("Part 1: %d\n", hor*depth)
}

func Part2() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	r, err := regexp.Compile(`(?P<Direction>[\w]+)\s(?P<Steps>[\d]+)`)
	var depth, hor, aim int
	for _, line := range lines {
		res := r.FindStringSubmatch(line) // res[0]: direction, res[1]: steps
		var direction = res[1]
		steps, _ := strconv.Atoi(res[2])
		switch direction {
		case "forward":
			hor += steps
			depth += aim * steps
		case "up":
			aim -= steps
		case "down":
			aim += steps
		default:
			panic(fmt.Sprintf("Unexpected direction value: %v", direction))
		}
	}
	fmt.Printf("Part 2: %d\n", hor*depth)
}

func main() {
	Part2()
}
