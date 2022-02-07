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

func Btoi(b bool) int64 {
	if b {
		return 1
	} else {
		return -1
	}
}

func Intersection(c1, c2 Cube) *Cube {
	if !(c1.minX <= c2.maxX && c1.maxX >= c2.minX) {
		return nil
	}
	if !(c1.minY <= c2.maxY && c1.maxY >= c2.minY) {
		return nil
	}
	if !(c1.minZ <= c2.maxZ && c1.maxZ >= c2.minZ) {
		return nil
	}
	minX := utils.Max(c1.minX, c2.minX)
	maxX := utils.Min(c1.maxX, c2.maxX)
	minY := utils.Max(c1.minY, c2.minY)
	maxY := utils.Min(c1.maxY, c2.maxY)
	minZ := utils.Max(c1.minZ, c2.minZ)
	maxZ := utils.Min(c1.maxZ, c2.maxZ)
	var action = Btoi(bool(c1.action)) * Btoi(bool(c2.action))
	if c1.action == c2.action {
		action = -Btoi(bool(c1.action))
	} else if c1.action == SetOn && c2.action == SetOff {
		action = 1
	}
	var bAction Action = false
	if action == 1 {
		bAction = true
	}
	return &Cube{
		action: bAction,
		minX:   minX,
		maxX:   maxX,
		minY:   minY,
		maxY:   maxY,
		minZ:   minZ,
		maxZ:   maxZ,
	}
}

func (c Cube) Volume() int64 {
	return int64((c.maxX - c.minX + 1) * (c.maxY - c.minY + 1) * (c.maxZ - c.minZ + 1))
}

func Part2() {
	lines, _ := utils.ReadLines("input.txt")
	cubes := make([]Cube, 0)
	for _, newCube := range ParseInput(lines) {
		is := make([]Cube, 0)
		for _, oldCube := range cubes {
			i := Intersection(newCube, oldCube)
			if i != nil {
				is = append(is, *i)
			}
		}
		for _, i := range is {
			cubes = append(cubes, i)
		}
		if newCube.action == SetOn {
			cubes = append(cubes, newCube)
		}
	}

	var num int64
	for _, cube := range cubes {
		volume := Btoi(bool(cube.action)) * cube.Volume()
		num += volume
	}
	fmt.Printf("Part 2 -> %v\n", num)
}

func Part1() {
	lines, _ := utils.ReadLines("input.txt")
	cubes := ParseInput(lines)
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
	fmt.Printf("Part 1 -> %d\n", num)
}

func main() {
	Part1()
	Part2()
}
