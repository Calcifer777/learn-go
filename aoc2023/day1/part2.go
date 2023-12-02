package day1

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func digits() map[string]int {
	return map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
}

func checkErr(err error) {
	if err != nil {
		slog.Error("Error: ", slog.Any("err", err))
		panic(err)
	}
}

func parseLine(line string) []int {
	nums := make([]int, 0)
	candidates := digits()
	for idx := range line {
		for s, i := range candidates {
			if idx+len(s) > len(line) {
				continue
			}
			if line[idx:idx+len(s)] == s {
				nums = append(nums, i)
			}
		}
	}
	return nums
}

// Part2 ...
func Part2() {
	path := "data/day1/full.txt"
	file, err := os.Open(path)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make([]int, 1000)
	for scanner.Scan() {
		nums := parseLine(scanner.Text())
		n := nums[0]*10 + nums[len(nums)-1]
		fmt.Printf("Line digits value is: %d\n", n)
		arr = append(arr, n)
	}

	total := 0
	for _, i := range arr {
		total += i
	}
	fmt.Printf("%v\n", total)
}
