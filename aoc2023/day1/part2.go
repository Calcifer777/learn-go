package day1

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func digits() map[int]string {
	return map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
}

func checkErr(err error) {
	if err != nil {
		slog.Error("Error: ", slog.Any("err", err))
		panic(err)
	}
}

func parseLine(s string) []int {
	fmt.Printf("Parsing %s\n", s)
	nums := make([]int, 0)
	candidates := digits()
	for _, ch := range s {
		if i, err := strconv.Atoi(string(ch)); err == nil {
			fmt.Printf("Found digit: %d\n", i)
			nums = append(nums, i)
			candidates = digits()
		} else {
			for i, s := range candidates {
				if []rune(s)[0] == ch {
					candidates[i] = s[1:]
					if len(s) == 1 {
						fmt.Printf("Found %d!\n", i)
						nums = append(nums, i)
						candidates[i] = digits()[i]
					}
				}
			}
		}
	}
	fmt.Println("")
	return nums
}

// Part2 ...
func Part2() {
	path := "data/day1/part2/full.txt"
	file, err := os.Open(path)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make([]int, 0)
	for scanner.Scan() {
		nums := parseLine(scanner.Text())
		arr = append(arr, nums[0]*10+nums[len(nums)-1])
	}

	// fmt.Printf("Nums:\n%v", arr)

	sum := 0
	for _, i := range arr {
		sum += i
	}
	fmt.Printf("The sum is: %v\n", sum)
}
