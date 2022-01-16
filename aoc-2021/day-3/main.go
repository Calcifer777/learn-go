package main

import (
	"fmt"
	"strconv"
	"utils"
)

func Part1() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	var counts = make([]int, len(lines[0]))
	for _, line := range lines {
		for j, bit := range line {
			if bit == '1' {
				counts[j] += 1
			}
		}
	}
	var size = len(lines)
	var strGamma string
	var strEpsilon string
	for _, x := range counts {
		if x >= size/2 {
			strGamma += "1"
			strEpsilon += "0"
		} else {
			strGamma += "0"
			strEpsilon += "1"
		}
	}
	gamma, _ := strconv.ParseInt(strGamma, 2, 64)
	epsilon, _ := strconv.ParseInt(strEpsilon, 2, 64)
	fmt.Printf("Part 1: %d\n", gamma*epsilon)
}

func Filter(vs []string, pos int, value byte) []string {
	filtered := make([]string, 0)
	for _, v := range vs {
		if v[pos] == value {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func MostCommon(lines []string, idx int) rune {
	var count int = 0
	for _, line := range lines {
		if line[idx] == '1' {
			count++
		}
	}
	var mostCommon rune
	if float64(count) >= float64(len(lines))/2 {
		mostCommon = '1'
	} else {
		mostCommon = '0'
	}
	return mostCommon
}

func Part2() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)

	// Find oxygen
	var oxygen []string = lines
	for idx := 0; len(oxygen) > 1; idx++ {
		mostCommon := MostCommon(oxygen, idx)
		switch mostCommon {
		case '1':
			oxygen = Filter(oxygen, idx, '1')
		case '0':
			oxygen = Filter(oxygen, idx, '0')
		}
	}
	iOxygen, _ := strconv.ParseInt(oxygen[0], 2, 64)
	fmt.Printf("Oxygen: %v -> %d\n", oxygen, iOxygen)

	// Find co2
	var co2 []string = lines
	for idx := 0; len(co2) > 1; idx++ {
		mostCommon := MostCommon(co2, idx)
		switch mostCommon {
		case '1':
			co2 = Filter(co2, idx, '0')
		case '0':
			co2 = Filter(co2, idx, '1')
		}
	}
	iCo2, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Printf("CO2: %v -> %d\n", co2, iCo2)

	fmt.Printf("Part 2: %d\n", iOxygen*iCo2)
}

func main() {
	Part1()
}
