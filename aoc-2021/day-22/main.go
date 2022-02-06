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

func Btoi(b bool) int {
	if b {
		return 1
	} else {
		return -1
	}
}

func Intersection(c1, c2 Cube) int64 {
  // fmt.Printf("%+v\n", c1)
  // fmt.Printf("%+v\n", c2)
	var x, y, z int64
	var fst, snd Cube
	// X
	if c1.minX < c2.minX {
		fst = c1
		snd = c2
	} else {
    fst = c2
    snd = c1
  }
	if fst.maxX < snd.minX {
		return 0
	} else if fst.minX <= snd.minX && fst.maxX >= snd.minX && fst.maxX <= snd.maxX {
		x = int64(fst.maxX - snd.minX + 1)
	} else if fst.maxX > snd.maxX {
		x = int64(snd.maxX - snd.minX + 1)
	}
	// Y
	if c1.minY < c2.minY {
		fst = c1
		snd = c2
	}
	if fst.maxY < snd.minY {
		return 0
	} else if fst.minY <= snd.minY && fst.maxY >= snd.minY && fst.maxY <= snd.maxY {
		y = int64(fst.maxY - snd.minY + 1)
	} else if fst.maxY > snd.maxY {
		y = int64(snd.maxY - snd.minY + 1)
	}
	// Z
	if c1.minZ < c2.minZ {
		fst = c1
		snd = c2
	}
	if fst.maxZ < snd.minZ {
		return 0
	} else if fst.minZ <= snd.minZ && fst.maxZ >= snd.minZ && fst.maxZ <= snd.maxZ {
		z = int64(fst.maxZ - snd.minZ + 1)
	} else if fst.maxZ > snd.maxZ {
		z = int64(snd.maxZ - snd.minZ + 1)
	}
	return int64(x * y * z)
}

func main() {
	lines, _ := utils.ReadLines("input-sample-3.txt")
	cubes := ParseInput(lines)
  // fmt.Printf("%d\n", Intersection(cubes[0], cubes[0]))
  intersections := make([][]int64, len(cubes))
  for i := 0; i < len(cubes); i++ {
    intersections[i] = make([]int64, len(cubes))
    for j := 0; j < len(cubes); j++ {
      intersections[i][j] = Intersection(cubes[i], cubes[j])
    }
  }
  fmt.Printf("%v\n", intersections)
}

func Part2() {
	lines, _ := utils.ReadLines("input-sample-1.txt")
	cubes := ParseInput(lines)
	var num int64
	for _, c := range cubes {
		fmt.Printf("%v\n", c)
		num += int64(Btoi(bool(c.action)) * (c.maxX - c.minX) * (c.maxY - c.minY) * (c.maxY - c.minY))
	}
	fmt.Printf("Num cubes on %d\n", num)
}

func Part1() {
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
