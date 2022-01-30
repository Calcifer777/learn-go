package main

import (
	"fmt"
	"strings"
	"utils"
)

type Pair struct {
	_1 string
	_2 string
}

type Mappings map[string]Pair
type MPoly map[string]int

// Parse the input
// Polymer: string
// Insertion rules: map[string][Pair(of strings)]
func ParseInput(lines []string) (string, Mappings) {
	seq := lines[0]
	mappings := make(Mappings)
	for _, l := range lines[2:] {
		chunks := strings.Split(l, " -> ")
		mappings[chunks[0]] = Pair{chunks[0][0:1] + chunks[1], chunks[1] + chunks[0][1:]}
	}
	return seq, mappings
}

// Yield a map of all the character pairs in a string
func MapPolymer(sPoly string, m Mappings) MPoly {
	mPolymer := make(MPoly)
	l := len(sPoly)
	for i := 0; i < l-1; i++ {
		_, ok := mPolymer[sPoly[i:i+2]]
		if ok {
			mPolymer[sPoly[i:i+2]] += 1
		} else {
			mPolymer[sPoly[i:i+2]] = 1
		}
	}
	return mPolymer
}

// Implement the insertion rule
// Leverage a second MPoly table to prevent insertions' effects
// to influence each other
func Insert(mPoly MPoly, m Mappings) MPoly {
	mPolyNew := make(MPoly)
	for k, v := range mPoly {
		if pair, isin := m[k]; isin {
			mPolyNew[pair._1] += v
			mPolyNew[pair._2] += v
		} else {
			mPolyNew[k] += v
		}
	}
	return mPolyNew
}

// Re-map each pair occurrence to their corresponding character frequencies,
// adjust for the first and last character frequencies (1/2 more times more each)
// Take the diff between max and min frequencies
func Score(mPoly MPoly, first string, last string) int {
	scores := make(MPoly)
	for k, v := range mPoly {
		scores[k[:1]] += v
		scores[k[1:]] += v
	}
	scores[first] += 1
	scores[last] += 1
	min := int(^uint(0) >> 1)
	max := 0
	for _, v := range scores {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return (max - min) / 2
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	sPoly, mappings := ParseInput(lines)
	mPoly := MapPolymer(sPoly, mappings)
	for i := 0; i < 10; i++ {
		mPoly = Insert(mPoly, mappings)
	}
	score := Score(mPoly, sPoly[:1], sPoly[len(sPoly)-1:])
	fmt.Printf("Part 1 -> %d\n", score)
	for i := 0; i < 30; i++ {
		mPoly = Insert(mPoly, mappings)
	}
	score = Score(mPoly, sPoly[:1], sPoly[len(sPoly)-1:])
	fmt.Printf("Part 2 -> %d\n", score)
}
