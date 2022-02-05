package main

import (
	"fmt"
	"math"
	"strings"
	"utils"
)

type Point struct{ x, y int }
type Image map[Point]bool

func (image Image) String() string {
	// Find image size
	var minX = math.MaxInt64
	var minY = math.MaxInt64
	var maxX = math.MinInt64
	var maxY = math.MinInt64
	for p := range image {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}
	rangeX := maxX - minX
	rangeY := maxY - minY
	// Create grid
	grid := make([][]string, rangeY+1)
	for y := 0; y <= rangeY; y++ {
		grid[y] = strings.Split(strings.Repeat(".", rangeX+1), "")
	}
	for p := range image {
		grid[p.y-minY][p.x-minX] = "#"
	}
	// Join into a single string
	str := ""
	for _, r := range grid {
		str += strings.Join(r, "") + "\n"
	}
	return str
}

type Boundary struct {
	minX, maxX, minY, maxY int
}

func (image Image) Boundary() Boundary {
	var minX = math.MaxInt64
	var minY = math.MaxInt64
	var maxX = math.MinInt64
	var maxY = math.MinInt64
	for p := range image {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}
	return Boundary{minX, maxX, minY, maxY}
}

func ParseImage(lines []string) Image {
	image := make(Image)
	for r, line := range lines {
		for c, char := range line {
			if char == '#' {
				image[Point{x: c, y: r}] = true
			}
		}
	}
	return image
}

func EnhanceOnce(image Image, algo string, unseenDefault bool) Image {
	enhanced := make(Image)
	b := image.Boundary()
	for x := b.minX - 2; x <= b.maxX+2; x++ {
		for y := b.minY - 2; y <= b.maxY+2; y++ {
			c := 0
			algoIdx := 0
			for yy := y + 1; yy >= y-1; yy-- {
				for xx := x + 1; xx >= x-1; xx-- {
					if xx > b.maxX || xx < b.minX || yy > b.maxY || yy < b.minY {
						if unseenDefault {
							algoIdx += (1 << c)
						}
					} else if ok := image[Point{xx, yy}]; ok {
						algoIdx += (1 << c)
					}
					c++
				}
			}
			if algo[algoIdx] == '#' {
				enhanced[Point{x, y}] = true
			}
		}
	}
	return enhanced
}

func UnseenOn(start bool, algo string) bool {
	var idx int = 0
	if start {
		idx = 511
	}
	if algo[idx] == '#' {
		return true
	} else {
		return false
	}
}

func Enhance(image Image, algo string, n int) Image {
	unseenDefault := false
	for i := 1; i <= n; i++ {
		image = EnhanceOnce(image, algo, unseenDefault)
		unseenDefault = UnseenOn(unseenDefault, algo)
	}
	return image
}

func Part1() {
	lines, _ := utils.ReadLines("input.txt")
	algo := lines[0]
	image := ParseImage(lines[2:])
	image = Enhance(image, algo, 2)
	fmt.Printf("Part 1 -> %d\n", len(image))
}

func Part2() {
	lines, _ := utils.ReadLines("input.txt")
	algo := lines[0]
	image := ParseImage(lines[2:])
	image = Enhance(image, algo, 50)
	fmt.Printf("Part 2 -> %d\n", len(image))
}

func main() {
	Part1()
	Part2()
}
