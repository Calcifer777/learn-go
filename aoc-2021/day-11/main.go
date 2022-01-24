package main

import (
	"fmt"
	"strconv"
	"utils"
)

type Octo struct {
	value   int
	flashed bool
}

type Swarm [][]*Octo

func ParseInput(lines []string) Swarm {
	octos := make(Swarm, 0)
	for _, line := range lines {
		row := make([]*Octo, 0)
		for _, c := range line {
			i, _ := strconv.Atoi(string(c))
			row = append(row, &Octo{i, false})
		}
		octos = append(octos, row)
	}
	return octos
}

func (o *Octo) String() string {
	return fmt.Sprintf("%d", o.value)
}

func (o *Octo) Activate() {
	o.value = (o.value + 1) % 10
	return
}

func (s *Swarm) String() string {
	var str = ""
	for _, row := range *s {
		for _, o := range row {
			str += fmt.Sprintf("%d,", o.value)
		}
		str += "\n"
	}
	return str
}

func (s *Swarm) Activate(minR, maxR, minC, maxC int) int {
	swarm := *s
	flashes := 0
	for ir := minR; ir <= maxR; ir++ {
		for ic := minC; ic <= maxC; ic++ {
			if swarm[ir][ic].flashed {
				continue
			}
			swarm[ir][ic].Activate()
			if swarm[ir][ic].value == 0 {
				swarm[ir][ic].flashed = true
				flashes += 1 +
					s.Activate(
						utils.Max(ir-1, 0),
						utils.Min(ir+1, len(swarm)-1),
						utils.Max(ic-1, 0),
						utils.Min(ic+1, len(swarm[0])-1),
					)
			}
		}
	}
	return flashes
}

func (s *Swarm) Syncd() bool {
	swarm := *s
	for ir := 0; ir < len(swarm); ir++ {
		for ic := 0; ic < len(swarm[0]); ic++ {
			if swarm[ir][ic].value > 0 {
        return false
      }
		}
	}
  return true
}

func (s *Swarm) IterActivate(n int, flagSyncd bool) int {
  flashes := 0
  swarm := *s
	height := len(swarm)
	width := len(swarm[0])
  for i := 0; i < n; i++ {
    flashes += swarm.Activate(0, height-1, 0, width-1)
    swarm.Reset()
    if swarm.Syncd() {
      fmt.Printf("Swarm syncd at step %d\n", i+1)
      if flagSyncd {
        return flashes
      }
    }
		// fmt.Printf("%s\n", swarm.String())
    // fmt.Printf("Flashes after %d steps: %d\n", i, flashes)
  }
  return flashes
}

func (s *Swarm) Reset() {
	swarm := *s
	for ir := 0; ir < len(swarm); ir++ {
		for ic := 0; ic < len(swarm[0]); ic++ {
			swarm[ir][ic].flashed = false
		}
	}
}

func Part1() {
	lines, _ := utils.ReadLines("input.txt")
	swarm := ParseInput(lines)
  flashes := swarm.IterActivate(100, false)
  fmt.Printf("Part 1 -> %d\n", flashes)
}

func Part2() {
	lines, _ := utils.ReadLines("input.txt")
	swarm := ParseInput(lines)
  swarm.IterActivate(1000, true)
}

func main() {
  Part2()
}
