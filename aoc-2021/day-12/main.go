package main

import (
	"fmt"
	"strings"
	"utils"
)

type Size int64

const (
	Small Size = 0
	Big        = 1
)

type Cave struct {
	name  string
	links []*Cave
	size  Size
}

type CaveSystem map[string]*Cave

func ParseInput(lines []string) CaveSystem {
	caves := make(map[string]*Cave, 0)
	// Parse caves
	for _, line := range lines {
		cns := strings.Split(line, "-")
		for _, cn := range cns {
			if _, ok := caves[cn]; !ok {
				var size Size
				if cn == strings.ToLower(cn) {
					size = Small
				} else {
					size = Big
				}
				caves[cn] = &Cave{cn, make([]*Cave, 0), size}
			}
		}
		cnFrom := cns[0]
		cnTo := cns[1]
		if caveFrom, ok := caves[cnFrom]; ok {
			caveFrom.links = append(caveFrom.links, caves[cnTo])
			caves[cnFrom] = caveFrom
		}
		if caveTo, ok := caves[cnTo]; ok {
			caveTo.links = append(caveTo.links, caves[cnFrom])
			caves[cnTo] = caveTo
		}
	}
	return caves
}

func Map[T any, V any](ts []T, f func(t T) V) []V {
	vs := make([]V, 0)
	for _, t := range ts {
		vs = append(vs, f(t))
	}
	return vs
}

func IsIn[T comparable](arr []T, t T) bool {
	for _, v := range arr {
		if v == t {
			return true
		}
	}
	return false
}

func Explore(cave *Cave) int {
	var loop func(cave *Cave, explored []*Cave) int
	loop = func(cave *Cave, explored []*Cave) int {
		paths := 0
		for _, c := range cave.links {
			if c.name == "end" {
				paths += 1
			} else if c.size == Small && IsIn(explored, c) {
				continue
			} else if c.size == Small && !IsIn(explored, c) {
				paths += loop(c, append(explored, c))
			} else {
				paths += loop(c, explored)
			}
		}
		return paths
	}
	return loop(cave, []*Cave{cave})
}

func ExploreTwice(cave *Cave) int {
	var loop func(cave *Cave, explored []*Cave, exploredTwice bool) int
	loop = func(cave *Cave, explored []*Cave, exploredTwice bool) int {
		paths := 0
		for _, c := range cave.links {
			if c.name == "end" {
				paths += 1
			} else if c.size == Big {
				paths += loop(c, explored, exploredTwice)
			} else if c.size == Small {
				if !IsIn(explored, c) {
					paths += loop(c, append(explored, c), exploredTwice)
				} else {
					if !exploredTwice && c.name != "start" {
						paths += loop(c, append(explored, c), true)
					} else {
						continue
					}
				}
			}
		}
		return paths
	}
	return loop(cave, []*Cave{cave}, false)
}

func Part1() {
	lines, _ := utils.ReadLines("input.txt")
	caves := ParseInput(lines)
	paths := Explore(caves["start"])
	fmt.Printf("Part 1 -> %d\n", paths)
}

func Part2() {
	lines, _ := utils.ReadLines("input.txt")
	caves := ParseInput(lines)
	paths := ExploreTwice(caves["start"])
	fmt.Printf("Part 2 -> %d\n", paths)
}

func main() {
	Part1()
	Part2()
}
