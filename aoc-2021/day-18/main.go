package main

import (
	"fmt"
	"strconv"
	"utils"
)

type P struct {
	value int
	level int
}

func (p P) String() string {
  return fmt.Sprintf("P{%d,%d}", p.value, p.level)
}

func ParseInput(s string) []P {
	pairs := make([]P, 0)
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
				pairs = append(pairs, P{x, open})
				n = ""
			}
			if c == ']' {
				open--
			}
		}
	}
	return pairs
}

func Add(a, b []P) []P {
	inc := func(p P) P { return P{p.value, p.level + 1} }
	a1 := utils.Map(a, inc)
	b1 := utils.Map(b, inc)
	return append(a1, b1...)
}

func Explode(ps []P) []P {
  xs := make([]P, len(ps))
  copy(xs, ps)
	// Explode pairs
	var explodedIdx int = -1
	for i := 0; i < len(xs); i++ {
		if xs[i].level > 4 {
			if i == 0 {
				xs[i+2] = P{xs[i+2].value + xs[i+1].value, xs[i+2].level}
				xs[i] = P{0, xs[i].level}
        explodedIdx = 1
			} else if i == len(xs)-2 {
				xs[i-1] = P{xs[i-1].value + xs[i].value, xs[i-1].level}
				xs[i+1] = P{0, xs[i+1].level}
				explodedIdx = i
			} else if xs[i].level == xs[i+1].level {
				xs[i-1] = P{xs[i-1].value + xs[i].value, xs[i-1].level}
				xs[i+2] = P{xs[i+2].value + xs[i+1].value, xs[i+2].level}
				xs[i+1] = P{0, xs[i+1].level}
				explodedIdx = i
			}
			break
		}
	}
  if explodedIdx == -1 {
    return xs
  }
  // fmt.Printf("%v\n", xs)
  // fmt.Printf("Exploded %v\n", explodedIdx)
	// Adjust levels
	for i := explodedIdx+1; i < len(xs); i++ {
		if xs[i].level >= xs[explodedIdx].level {
			xs[i].level--
		} else {
      break
    }
	}
	for i := explodedIdx-1; i >= 0; i-- {
		if xs[i].level >= xs[explodedIdx].level {
      // fmt.Printf("%v\n", xs[i])
			xs[i].level--
		} else {
      break
    }
	}
	// Drop exploded Pair
	return append(xs[:explodedIdx], xs[explodedIdx+1:]...)
}

func main() {
	// s1 := "[[[[[9,8],1],2],3],4]" // [[[[0,9],2],3],4]
	// s1 := "[7,[6,[5,[4,[3,2]]]]]" // [7,[6,[5,[7,0]]]]
	// s1 := "[[6,[5,[4,[3,2]]]],1]" // [[6,[5,[7,0]]],3]
	// s1 := "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]" // [[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]
	s1 := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]" // [[3,[2,[8,0]]],[9,[5,[7,0]]]]
	xs := ParseInput(s1)
	fmt.Printf("%v\n", xs)
	fmt.Printf("%v\n", Explode(xs))
	// s2 := "[1, [2, [7, 0]]]"
	// p1 := ParseInput(s1)
	// p2 := ParseInput(s2)
	// fmt.Printf("%+v\n", p1)
	// fmt.Printf("%+v\n", p2)
	// fmt.Printf("%+v\n", Add(p1, p2))
}
