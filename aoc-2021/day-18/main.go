package main

import (
	"fmt"
	"reflect"
	"strconv"
	"utils"
)

type N struct {
	value int
	level int
}

func (p N) String() string {
	return fmt.Sprintf("N{%d,%d}", p.value, p.level)
}

func ParseN(s string) []N {
	pairs := make([]N, 0)
	var open int
	var n string
	for _, c := range s {
		switch c {
		case '[':
			open++
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n += string(c)
		case ',', ']':
			if n != "" {
				x, _ := strconv.Atoi(n)
				pairs = append(pairs, N{x, open})
				n = ""
			}
			if c == ']' {
				open--
			}
		}
	}
	return pairs
}

func Add(a, b []N) []N {
	inc := func(p N) N { return N{p.value, p.level + 1} }
	a1 := utils.Map(a, inc)
	b1 := utils.Map(b, inc)
	return append(a1, b1...)
}

func Explode(ps []N) []N {
	xs := make([]N, len(ps))
	copy(xs, ps)
	// Explode pairs
	var explodedIdx int = -1
	for i := 0; i < len(xs); i++ {
		if xs[i].level > 4 {
			if i == 0 {
				xs[i+2] = N{xs[i+2].value + xs[i+1].value, xs[i+2].level}
				xs[i] = N{0, xs[i].level - 1}
				explodedIdx = 1
			} else if i == len(xs)-2 {
				xs[i-1] = N{xs[i-1].value + xs[i].value, xs[i-1].level}
				xs[i+1] = N{0, xs[i+1].level - 1}
				explodedIdx = i
			} else if xs[i].level == xs[i+1].level {
				xs[i-1] = N{xs[i-1].value + xs[i].value, xs[i-1].level}
				xs[i+2] = N{xs[i+2].value + xs[i+1].value, xs[i+2].level}
				xs[i+1] = N{0, xs[i+1].level - 1}
				explodedIdx = i
			}
			break
		}
	}
	// If no exploded pair, return original array
	if explodedIdx == -1 {
		return xs
	} else {
		// Drop exploded pair
		return append(xs[:explodedIdx], xs[explodedIdx+1:]...)
	}
}

func Split(ps []N) []N {
	xs := make([]N, len(ps))
	copy(xs, ps)
	var splitIdx int = -1
	for idx, p := range xs {
		if p.value >= 10 {
			splitIdx = idx
			break
		}
	}
	if splitIdx == -1 {
		return xs
	}
	newPairs := []N{
		N{xs[splitIdx].value / 2, xs[splitIdx].level + 1},
		N{(xs[splitIdx].value + 1) / 2, xs[splitIdx].level + 1},
	}
	splitted := append(xs[:splitIdx], newPairs...)
	splitted = append(splitted, ps[splitIdx+1:]...)
	return splitted
}

func ReduceOnce(ps []N) []N {
	xs := make([]N, len(ps))
	copy(xs, ps)
	// fmt.Printf("exploding...")
	exploded := Explode(xs)
	if !reflect.DeepEqual(exploded, ps) {
		// fmt.Printf("ok\n")
		return exploded
	} else {
		// fmt.Printf("nothing to do, splitting...\n")
		return Split(exploded)
	}
}

func Reduce(ps []N) []N {
	curr := make([]N, len(ps))
	copy(curr, ps)
	prev := make([]N, len(ps))
	copy(prev, ps)
	var i int
	// fmt.Printf("Reducing:\n%v\n", ps)
	for {
		// fmt.Printf("Step %d: ", i)
		curr = ReduceOnce(curr)
		if reflect.DeepEqual(prev, curr) {
			// fmt.Printf("\tDone after %d steps\n", i-1)
			return curr
		}
		prev = make([]N, len(curr))
		copy(prev, curr)
		i++
	}
}

func ReduceList(xs [][]N) []N {
	if len(xs) == 0 {
		return make([]N, 0)
	} else if len(xs) == 1 {
		return xs[0]
	} else {
		sum := xs[0]
		for _, x := range xs[1:] {
			sum = Add(sum, x)
			sum = Reduce(sum)
		}
		return sum
	}
}

func Magnitude(ps []N) int {
	xs := make([]N, len(ps))
	copy(xs, ps)
	var flag bool
	for {
		flag = true
		for i := 0; i < len(xs)-1; i++ {
			if xs[i].level == xs[i+1].level {
				xs = append(
					append(xs[:i], N{3*xs[i].value + 2*xs[i+1].value, xs[i].level - 1}),
					xs[i+2:]...,
				)
				flag = false
				break
			}
		}
		if flag {
			break
		}
	}
	return xs[0].value
}

func LargestPairMagnitude(xs [][]N) int {
	var largest int
	for i := 0; i < len(xs); i++ {
		for j := 0; j < len(xs); j++ {
			if i == j {
				continue
			}
			sum := Add(xs[i], xs[j])
			reduced := Reduce(sum)
			m := Magnitude(reduced)
			if m > largest {
				largest = m
			}
		}
	}
	return largest
}

func Part1() {
	lines, _ := utils.ReadLines("input.txt")
	xs := utils.Map(lines, ParseN)
	result := ReduceList(xs)
	fmt.Printf("Part 1 -> %d\n", Magnitude(result))
}

func Part2() {
	lines, _ := utils.ReadLines("input.txt")
	xs := utils.Map(lines, ParseN)
	result := LargestPairMagnitude(xs)
	fmt.Printf("Part 2 -> %d\n", result)
}

func main() {
	Part1()
	Part2()
}
