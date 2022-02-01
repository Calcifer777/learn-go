package main

// https://adventofcode.com/2021/day/17

import (
	"fmt"
	"regexp"
	"strconv"
	"utils"
)

type Point struct{ x, y, vx, vy int }
type Area struct{ x0, x1, y0, y1 int }

func ParseInput(s string) Area {
	p := `target area: ` +
		`x=(?P<x0>-?\d+)..(?P<x1>-?\d+), ` +
		`y=(?P<y0>-?\d+)..(?P<y1>-?\d+)`
	r := regexp.MustCompile(p)
	matches := r.FindStringSubmatch(s)
	x0, _ := strconv.Atoi(matches[1])
	x1, _ := strconv.Atoi(matches[2])
	y0, _ := strconv.Atoi(matches[3])
	y1, _ := strconv.Atoi(matches[4])
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Area{x0, x1, y0, y1}
}

func (p Point) Next() Point {
	var vx int
	if p.vx > 0 {
		vx = p.vx - 1
	} else if p.vx < 0 {
		vx = p.vx + 1
	} else {
		vx = p.vx
	}
	return Point{
		x:  p.x + p.vx,
		y:  p.y + p.vy,
		vx: vx,
		vy: p.vy - 1,
	}
}

func IsIn(p Point, a Area) int {
	// Still too high
	if p.y > a.y1 {
		return -1
		// Hit
	} else if p.x >= a.x0 && p.x <= a.x1 && p.y >= a.y0 && p.y <= a.y1 {
		return 0
		// Within y, but off x
	} else if p.y < a.y0 {
		return 1
	} else {
		return -1
	}
}

func Shoot(p Point, a Area) int {
	i := 1
	maxY := p.y
	for {
		p = p.Next()
		if p.y > maxY {
			maxY = p.y
		}
		isin := IsIn(p, a)
		if isin == 0 {
			return maxY
		} else if isin == 1 {
			return 696969
		}
		i++
	}
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	area := ParseInput(lines[0])
	fmt.Printf("Area: %+v\n", area)
	gridSize := 1000
	grid := make([][]int, gridSize)
	for vx := 0; vx < gridSize; vx++ {
		grid[vx] = make([]int, gridSize)
		for vy := 0; vy < gridSize; vy++ {
			grid[vx][vy] = Shoot(Point{0, 0, vx - 500, vy - 500}, area)
		}
	}
	var bestVx, bestVy, maxY, distinct int
	for vx := 0; vx < gridSize; vx++ {
		for vy := 0; vy < gridSize; vy++ {
			if grid[vx][vy] != 696969 {
				distinct += 1
			}
			if grid[vx][vy] != 696969 && grid[vx][vy] > maxY {
				maxY = grid[vx][vy]
				bestVx = vx
				bestVy = vy
			}
		}
	}
	fmt.Printf("Best vx: %d, Best vy: %d, Max Y: %d, Distinct: %d\n", bestVx, bestVy, maxY, distinct)
	fmt.Printf("Part 1 -> %d\n", maxY)
	fmt.Printf("Part 2 -> %d\n", distinct)
}