package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"utils"
)

type V struct{ x, y, z int }
type AxisSorter []V
type BeaconsSet map[V]bool // use map[V]bool as set

func (b BeaconsSet) String() string {
	s := ""
	for v := range b {
		s += fmt.Sprintf("%v\n", v)
	}
	s += "\n"
	return s
}

func (b BeaconsSet) Keys() []V {
	keys := make([]V, len(b))
	idx := 0
	for k := range b {
		keys[idx] = k
		idx++
	}
	return keys
}

func (b BeaconsSet) Copy() BeaconsSet {
	newMap := make(BeaconsSet)
	for k, v := range b {
		newMap[k] = v
	}
	return newMap
}

const cos90 = 0
const sin90 = 1

func ParseInput(lines []string) map[int]BeaconsSet {
	input := make(map[int]BeaconsSet)
	patternH := regexp.MustCompile(`--- scanner (?P<section>\d+) ---`)
	patternV := regexp.MustCompile(`(?P<x>-?\d+),(?P<y>-?\d+),(?P<z>-?\d+)`)
	var section int
	var beacons BeaconsSet
	var sectionStart bool = true
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			sectionStart = true
			input[section] = beacons
			continue
		}
		if sectionStart {
			sectionStart = false
			matches := patternH.FindStringSubmatch(lines[i])
			s, _ := strconv.Atoi(matches[1])
			section = s
			beacons = make(BeaconsSet)
		} else {
			matches := patternV.FindStringSubmatch(lines[i])
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			z, _ := strconv.Atoi(matches[3])
			beacons[V{x, y, z}] = true
		}
	}
	input[section] = beacons
	return input
}

func (xs AxisSorter) Len() int      { return len(xs) }
func (xs AxisSorter) Swap(i, j int) { xs[i], xs[j] = xs[j], xs[i] }
func (xs AxisSorter) Less(i, j int) bool {
	if xs[i].x < xs[j].x {
		return true
	} else if xs[i].x > xs[j].x {
		return false
	} else {
		if xs[i].y < xs[j].y {
			return true
		} else if xs[i].y > xs[j].y {
			return false
		} else {
			if xs[i].z < xs[j].z {
				return true
			} else {
				return false
			}
		}
	}
}

func (v1 V) Sub(v2 V) V {
	return V{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

func (v1 V) Add(v2 V) V {
	return V{v1.x + v2.x, v1.y + v2.y, v1.z + v2.z}
}

func (v1 V) Mul(i int) V {
	return V{i * v1.x, i * v1.y, i * v1.z}
}

func Diff(xs, ys []V) []V {
	zs := make([]V, len(xs))
	for i := 0; i < len(xs); i++ {
		zs[i] = xs[i].Sub(ys[i])
	}
	return zs
}

// Rotate around X axis by 90 degrees counterclockwise
// Reference: https://en.wikipedia.org/wiki/Rotation_matrix
func RotateX(v V) V {
	return V{v.x, -sin90 * v.z, sin90 * v.y}
}

func RotateY(v V) V {
	return V{sin90 * v.z, v.y, -sin90 * v.x}
}

func RotateZ(v V) V {
	return V{-sin90 * v.y, sin90 * v.x, v.z}
}

func Rotate(v V, nx int, ny int, nz int) V {
	out := v
	for i := 0; i < nx; i++ {
		out = RotateX(out)
	}
	for i := 0; i < ny; i++ {
		out = RotateY(out)
	}
	for i := 0; i < nz; i++ {
		out = RotateZ(out)
	}
	return out
}

func FindOffset(xs []V, ys []V, threshold int) (V, bool) {
	diffMap := make(map[V]int)
	for _, x := range xs {
		diffMap[x.Sub(ys[0])] += 1
		for _, y := range ys[1:] {
			diffMap[x.Sub(y)] += 1
		}
	}
	for k, v := range diffMap {
		if v >= threshold {
			return k, true
		}
	}
	return *new(V), false
}

func FindOffsetWithRotation(xxs, yys BeaconsSet) (V, V, bool) {
	xs := xxs.Keys()
	ys := yys.Keys()
	sort.Sort(AxisSorter(xs))
	for nx := 0; nx < 4; nx++ {
		for ny := 0; ny < 4; ny++ {
			for nz := 0; nz < 4; nz++ {
				ysR := utils.Map(ys, func(v V) V { return Rotate(v, nx, ny, nz) })
				offset, found := FindOffset(xs, ysR, 12)
				if found {
					rotation := V{nx, ny, nz}
					// fmt.Printf("Offset found: %v, rot: %d\n", offset, rotation)
					return offset, rotation, true
				}
			}
		}
	}
	return *new(V), *new(V), false
}

type Mapping struct {
	offset V
	nx     int
	ny     int
	nz     int
	prev   int
}

func Merge(m map[int]BeaconsSet) (BeaconsSet, map[int]V) {
	done := map[int]bool{0: true}
	left := make(map[int]bool)
	for k := range m {
		left[k] = true
	}
	delete(left, 0)
	mappings := make(map[int]Mapping)
	offsets := make(map[int]V)
	offsets[0] = V{0, 0, 0}
	acc := m[0]
	for {
		if len(left) == 0 {
			break
		}
	outer:
		for idLeft := range left {
			for idDone := range done {
				// fmt.Printf("Comparing ids: %d - %d\n", idDone, idLeft)
				offset, rotation, ok := FindOffsetWithRotation(m[idDone], m[idLeft])
				if ok {
					mappings[idLeft] = Mapping{offset, rotation.x, rotation.y, rotation.z, idDone}
					done[idLeft] = true
					delete(left, idLeft)
					//
					currId := idLeft
					mapped := m[idLeft].Keys()
					off := V{0, 0, 0}
					for {
						mpng := mappings[currId]
						// fmt.Printf("Mapping %d to %d\n", currId, mpng.prev)
						mapped = utils.Map(mapped, func(v V) V { return Rotate(v, mpng.nx, mpng.ny, mpng.nz) })
						mapped = utils.Map(mapped, func(v V) V { return v.Add(mpng.offset) })
						off = off.Add(mpng.offset)
						if mpng.prev == 0 {
							break
						} else {
							currId = mpng.prev
						}
					}
					for _, k := range mapped {
						acc[k] = true
					}
					offsets[idLeft] = off
					break outer
				}
			}
		}
	}
	return acc, offsets
}

func LargestDistance(offsets map[int]V) int {
	var largest int
	for i := 0; i < len(offsets); i++ {
		for j := 0; j < len(offsets); j++ {
			o1 := offsets[i]
			o2 := offsets[j]
			dist := utils.Abs(o1.x-o2.x) + utils.Abs(o1.y-o2.y) + utils.Abs(o1.z-o2.z)
			if dist > largest {
				largest = dist
			}
		}
	}
	return largest
}

func Part1() {
	lines, _ := utils.ReadLines("input.txt")
	input := ParseInput(lines)
	beacons, _ := Merge(input)
	fmt.Printf("Part 1 -> %v\n", len(beacons))
}

func Part2() {
	lines, _ := utils.ReadLines("input.txt")
	input := ParseInput(lines)
	_, offsets := Merge(input)
	fmt.Printf("Part 2 -> %d\n", LargestDistance(offsets))  // 9634, sometimes is bugged, idk why
}

func main() {
	Part1()
	Part2()
}
