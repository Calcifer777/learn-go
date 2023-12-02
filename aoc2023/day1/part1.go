package day1

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

// Part1 ...
func Part1(path string) int {
	file, err := os.Open(path)
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
		}
		slog.Debug("Nums: ", slog.Any("nums", nums))
		arr = append(arr, nums[0]*10+nums[len(nums)-1])
	}
	sum := 0
	for _, i := range arr {
		sum += i
	}
	slog.Debug(fmt.Sprintf("The sum is: %v\n", sum))
	return sum
}
