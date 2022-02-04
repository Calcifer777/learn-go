package main

import (
	"fmt"
	"reflect"
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
				xs[i] = P{0, xs[i].level - 1}
				explodedIdx = 1
			} else if i == len(xs)-2 {
				xs[i-1] = P{xs[i-1].value + xs[i].value, xs[i-1].level}
				xs[i+1] = P{0, xs[i+1].level - 1}
				explodedIdx = i
			} else if xs[i].level == xs[i+1].level {
				xs[i-1] = P{xs[i-1].value + xs[i].value, xs[i-1].level}
				xs[i+2] = P{xs[i+2].value + xs[i+1].value, xs[i+2].level}
				xs[i+1] = P{0, xs[i+1].level - 1}
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

func Split(ps []P) []P {
	xs := make([]P, len(ps))
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
	newPairs := []P{
		P{xs[splitIdx].value / 2, xs[splitIdx].level + 1},
		P{(xs[splitIdx].value + 1) / 2, xs[splitIdx].level + 1},
	}
	splitted := append(xs[:splitIdx], newPairs...)
	splitted = append(splitted, ps[splitIdx+1:]...)
	return splitted
}

func ReduceOnce(ps []P) []P {
	xs := make([]P, len(ps))
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

func Reduce(ps []P) []P {
	curr := make([]P, len(ps))
	copy(curr, ps)
	prev := make([]P, len(ps))
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
		prev = make([]P, len(curr))
		copy(prev, curr)
		i++
	}
}

func ReduceList(xs [][]P) []P {
	if len(xs) == 0 {
		return make([]P, 0)
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

func Magnitude(ps []P) int {
	xs := make([]P, len(ps))
	copy(xs, ps)
  var flag bool
  for {
    flag = true
    for i:=0; i<len(xs)-1; i++ {
      if xs[i].level == xs[i+1].level {
        xs = append(
          append(xs[:i], P{3*xs[i].value+2*xs[i+1].value, xs[i].level-1}),
          xs[i+2:]...
        )
        flag = false
        break
      }
    }
    if flag { break }
  }
  return xs[0].value
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	xs := utils.Map(lines, ParseInput)
  result := ReduceList(xs)
	fmt.Printf("%v\n", result)
  fmt.Printf("Magnitude: %d\n", Magnitude(result))
}
