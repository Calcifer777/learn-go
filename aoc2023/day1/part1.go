package day1

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

// Part1 ...
func Part1() {
	file, err := os.Open("data/day1/part1/full.txt")
	if err != nil {
		slog.Error("Error in reading file", slog.Any("err", err))
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make([]int, 0)
	for scanner.Scan() {
		nums := make([]int, 0)
		for _, ch := range scanner.Text() {
			if i, err := strconv.Atoi(string(ch)); err == nil {
				nums = append(nums, i)
			}
			// fmt.Printf("%v\n", nums)
		}
		arr = append(arr, nums[0]*10+nums[len(nums)-1])
	}
	sum := 0
	for _, i := range arr {
		sum += i
	}
	fmt.Printf("The sum is: %v\n", sum)
}
