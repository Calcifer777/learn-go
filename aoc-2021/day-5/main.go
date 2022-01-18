package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

type HeatMap map[Point]int

func ParseInputLine(s string) Line {
	chunks := strings.Split(s, " -> ")
	if len(chunks) != 2 {
		panic(fmt.Sprintf("String %s cannot be parsed", s))
	}
	p1 := strings.Split(chunks[0], ",")
	p2 := strings.Split(chunks[1], ",")
	p1x, _ := strconv.Atoi(p1[0])
	p1y, _ := strconv.Atoi(p1[1])
	p2x, _ := strconv.Atoi(p2[0])
	p2y, _ := strconv.Atoi(p2[1])
	return Line{Point{p1x, p1y}, Point{p2x, p2y}}
}

func ParseInput(input []string) []Line {
	lines := make([]Line, 0)
	for _, s := range input {
		lines = append(lines, ParseInputLine(s))
	}
	return lines
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func (l Line) HVPoints() []Point {
	if l.start.x != l.end.x && l.start.y != l.end.y {
		return make([]Point, 0)
	}
	points := make([]Point, 0)
	if l.start.x == l.end.x {
		minY := min(l.start.y, l.end.y)
		for i := 0; i <= abs(l.start.y-l.end.y); i++ {
			points = append(points, Point{l.start.x, minY + i})
		}
	} else if l.start.y == l.end.y {
		minX := min(l.start.x, l.end.x)
		for i := 0; i <= abs(l.start.x-l.end.x); i++ {
			points = append(points, Point{minX + i, l.start.y})
		}
	}
	return points
}

func (l Line) Points() []Point {
	points := make([]Point, 0)
	if l.start.x == l.end.y && l.end.x == l.start.y {
		// 3,3 -> 7,7
		minX := min(l.start.x, l.end.x)
		maxX := max(l.start.x, l.end.x)
		for i := 0; i <= abs(l.start.x-l.end.x); i++ {
			points = append(points, Point{minX + i, maxX - i})
		}
	} else if l.start.x == l.end.x {
		// 3,3 -> 3,7
		minY := min(l.start.y, l.end.y)
		for i := 0; i <= abs(l.start.y-l.end.y); i++ {
			points = append(points, Point{l.start.x, minY + i})
		}
	} else if l.start.y == l.end.y {
		// 4,3 -> 7,3
		minX := min(l.start.x, l.end.x)
		for i := 0; i <= abs(l.start.x-l.end.x); i++ {
			points = append(points, Point{minX + i, l.start.y})
		}
	} else if (l.start.x-l.end.x)/(l.start.y-l.end.y) == 1 {
		// 2,4 -> 5,7
		minX := min(l.start.x, l.end.x)
		minY := min(l.start.y, l.end.y)
		for i := 0; i <= abs(l.start.x-l.end.x); i++ {
			points = append(points, Point{minX + i, minY + i})
		}
	} else if (l.start.x-l.end.x)/(l.start.y-l.end.y) == -1 {
		// 5,5 -> 8,2
		minX := min(l.start.x, l.end.x)
		maxY := max(l.start.y, l.end.y)
		for i := 0; i <= abs(l.start.x-l.end.x); i++ {
			points = append(points, Point{minX + i, maxY - i})
		}
	}
	return points
}

func NewHeatMap(points []Point) HeatMap {
	hm := make(map[Point]int)
	for _, p := range points {
		hm[p] += 1
	}
	return hm
}

func (hm *HeatMap) DangerZone() []Point {
	points := make([]Point, 0)
	for k, v := range *hm {
		if v >= 2 {
			points = append(points, k)
		}
	}
	return points
}

func Part1() {
	input, err := utils.ReadLines("input.txt")
	utils.Check(err)
	// Parse Lines
	lines := make([]Line, 0)
	for _, s := range input {
		lines = append(lines, ParseInputLine(s))
	}
	// Extract points
	points := make([]Point, 0)
	for _, l := range lines {
		points = append(points, l.HVPoints()...)
	}
	hm := NewHeatMap(points)
	dangerZone := hm.DangerZone()
	fmt.Printf("Part 1 -> Danger Zone size: %d\n", len(dangerZone))

}

func Part2() {
	input, err := utils.ReadLines("input.txt")
	utils.Check(err)
	// Parse Lines
	lines := make([]Line, 0)
	for _, s := range input {
		lines = append(lines, ParseInputLine(s))
	}
	// Extract points
	points := make([]Point, 0)
	for _, l := range lines {
		points = append(points, l.Points()...)
	}
	hm := NewHeatMap(points)
	dangerZone := hm.DangerZone()
	fmt.Printf("Part 2 -> Danger Zone size: %d\n", len(dangerZone))
}

func main() {
	Part1()
	Part2()
}
