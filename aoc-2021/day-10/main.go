package main

import (
	"fmt"
	"sort"
	"utils"
)

var OPEN = []rune{'(', '[', '{', '<'}
var OPEN_TO_CLOSED = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}
var CLOSED_TO_SCORE = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func IsIn[T comparable](arr []T, value T) bool {
	for _, t := range arr {
		if t == value {
			return true
		}
	}
	return false
}

func CheckCorrupted(line string) (int, bool) {
	stack := make([]rune, 0)
	for i, c := range line {
		if len(stack) == 0 && !IsIn(OPEN, c) {
			return i, true
		} else if IsIn(OPEN, c) {
			stack = append(stack, c)
		} else if !IsIn(OPEN, c) {
			if c == OPEN_TO_CLOSED[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return i, true
			}
		} else {
			fmt.Printf("Unreachable: stack %v, value: %v\n", stack, string(c))
		}
	}
	return -1, false
}

func Reverse(xs []rune) []rune {
	r := make([]rune, len(xs))
	for i := 0; i < len(xs); i++ {
		r[i] = xs[len(xs)-i-1]
	}
	return r
}

func AutoComplete(line string) ([]rune, bool) {
	stack := make([]rune, 0)
	for _, c := range line {
		if len(stack) == 0 && !IsIn(OPEN, c) {
			return make([]rune, 0), false
		} else if IsIn(OPEN, c) {
			stack = append(stack, c)
		} else if !IsIn(OPEN, c) {
			if c == OPEN_TO_CLOSED[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return make([]rune, 0), false
			}
		} else {
			fmt.Printf("Unreachable: stack %v, value: %v\n", stack, string(c))
		}
	}
	if len(stack) == 0 {
		return make([]rune, 0), true
	}
	ac := make([]rune, 0)
	for _, c := range Reverse(stack) {
		ac = append(ac, OPEN_TO_CLOSED[c])
	}
	return ac, true
}

func ScoreCompletion(arr []rune) int {
	score := 0
	m := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	for _, r := range arr {
		score *= 5
		score += m[r]
	}
	return score
}

func Part1() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	score := 0
	for _, line := range lines {
		i, ok := CheckCorrupted(line)
		if ok {
			score += CLOSED_TO_SCORE[rune(line[i])]
		}
	}
	fmt.Printf("Part 1: %d\n", score)
}

func Part2() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	scores := make([]int, 0)
	for _, line := range lines {
		_, corr := CheckCorrupted(line)
		if !corr {
			ac, _ := AutoComplete(line)
			scores = append(scores, ScoreCompletion(ac))
		}
	}
	sort.Ints(scores)
	fmt.Printf("Part 2: %d\n", scores[len(scores)/2])
}

func main() {
	Part1()
	Part2()
}
