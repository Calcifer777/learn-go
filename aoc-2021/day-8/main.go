package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"utils"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func ParseLine(s string) ([]string, []string) {
	chunks := strings.Split(s, "|")
	train := strings.Fields(chunks[0])
	test := strings.Fields(chunks[1])
	train_sorted := make([]string, 0)
	for _, s := range train {
		train_sorted = append(train_sorted, SortString(s))
	}
	test_sorted := make([]string, 0)
	for _, s := range test {
		test_sorted = append(test_sorted, SortString(s))
	}
	return train_sorted, test_sorted
}

func ContainsChars(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	flag := true
	for _, c := range s2 {
		if !strings.Contains(s1, string(c)) {
			return false
		}
	}
	return flag
}

func All(arr []string) bool {
	for _, x := range arr {
		if len(x) == 0 {
			return false
		}
	}
	return true
}

func Decode(xs []string) []string {
	m := make([]string, 10)
	for i := 0; i < 10; i++ {
		for _, s := range xs {
			l := len(s)
			switch l {
			case 2:
				m[1] = s
			case 3:
				m[7] = s
			case 4:
				m[4] = s
			case 7:
				m[8] = s
			case 5:
				if ContainsChars(m[9], s) && ContainsChars(s, m[1]) {
					m[3] = s
				} else if ContainsChars(m[9], s) {
					m[5] = s
				} else {
					m[2] = s
				}
			case 6:
				if ContainsChars(s, m[4]) {
					m[9] = s
				} else if ContainsChars(s, m[1]) {
					m[0] = s
				} else {
					m[6] = s
				}
			}
		}
	}
	return m
}

func Part1() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	sum := 0
	for _, line := range lines {
		_, test := ParseLine(line)
		for _, s := range test {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				sum += 1
			}
		}
	}
	fmt.Printf("Part 1 -> %d\n", sum)
}

func IndexOf(xs []string, s string) int {
	for idx, x := range xs {
		if x == s {
			return idx
		}
	}
	return -1
}

func ChunksToInt(chunks []string, mappings []string) int {
	m := make(map[string]int)
	for idx, s := range mappings {
		m[s] = idx
	}
	var value float64 = 0
	for idx, s := range chunks {
		value += float64(m[s]) * math.Pow(float64(10), float64(3-idx))
	}
	return int(value)
}

func Part2() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	var sum int
	for _, line := range lines {
		train, test := ParseLine(line)
		mappings := Decode(train)
		sum += ChunksToInt(test, mappings)
	}
	fmt.Printf("Part 2 -> %d\n", sum)
}

func main() {
	Part1()
	Part2()
}
