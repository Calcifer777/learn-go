package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func Part1() {
	lines, err := readLines("input.txt")
	check(err)
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
	lines, err := readLines("input.txt")
	check(err)
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
