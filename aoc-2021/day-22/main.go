package main

import (
	"fmt"
	"regexp"
	"strconv"
	"utils"
)

type Action bool

const (
	SetOn  Action = true
	SetOff Action = false
)

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type Cube struct {
	minX, maxX, minY, maxY, minZ, maxZ int
	action                             Action
}

func Between(x, min, max int) bool {
	return x >= min && x <= max
}

func ParseInput(lines []string) []Cube {
	cubes := make([]Cube, 0)
	p := regexp.MustCompile(`(?P<minX>on|off) x=(?P<minX>-?\d+)..(?P<maxX>-?\d+),y=(?P<minY>-?\d+)..(?P<maxY>-?\d+),z=(?P<minZ>-?\d+)..(?P<maxZ>-?\d+)`)
	for _, line := range lines {
		values := p.FindStringSubmatch(line)
		c := Cube{
			action: values[1] == "on",
			minX:   Atoi(values[2]),
			maxX:   Atoi(values[3]),
			minY:   Atoi(values[4]),
			maxY:   Atoi(values[5]),
			minZ:   Atoi(values[6]),
			maxZ:   Atoi(values[7]),
		}
		cubes = append(cubes, c)
	}
	return cubes
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	cubes := ParseInput(lines)
	for _, c := range cubes {
		fmt.Printf("%v\n", c)
	}
	num := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				isOn := false
				for _, c := range cubes {
					if Between(x, c.minX, c.maxX) && Between(y, c.minY, c.maxY) && Between(z, c.minZ, c.maxZ) {
						isOn = bool(c.action)
					}
				}
				if isOn {
					num++
				}
			}
		}
	}
	fmt.Printf("Num cubes on %d\n", num)
}
