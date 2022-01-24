package main

import (
  "fmt"
  "sort"
  "strconv"
  "utils"
)

type Heatmap [][]int

func ParseInput(input []string) Heatmap {
  hm := make(Heatmap, 0)
  for _, line := range input {
    row := make([]int, 0)
    for _, c := range line {
      i, _ := strconv.Atoi(string(c))
      row = append(row, i)
    }
    hm = append(hm, row)
  }
  return hm
}

type Point struct {
  x, y int
}

func (hm Heatmap) LowPoints() []Point {
  ws := 1
  heigth := len(hm)
  width := len(hm[0])
  lp := make([]Point, 0)
  for ir, r := range hm {
    for ic, v := range r {
      minR := utils.Max(ir-ws, 0)
      maxR := utils.Min(ir+ws, heigth-1)
      minC := utils.Max(ic-ws, 0)
      maxC := utils.Min(ic+ws, width-1)
      minV := 99
      for i:=minR; i <=maxR; i++ {
        for j:=minC; j <=maxC; j++ {
          if hm[i][j] < minV {
            minV = hm[i][j]
          }
        }
      }
      if v <= minV {
        lp = append(lp, Point{ir, ic})
      }
    }
  }
  return lp
}

func (hm Heatmap) Risk() int {
  lps := hm.LowPoints()
  risk := 0
  for _, p := range lps {
    risk += hm[p.x][p.y] + 1
  }
  return risk
}

func Part1() {
  input, _ := utils.ReadLines("input.txt")
  hm := ParseInput(input)
  fmt.Printf("Part 1: %v\n", hm.Risk())
}

func IsIn(points []Point, p Point) bool {
  for _, point := range points {
    if point == p {
      return true
    }
  }
  return false
}

func Basin(hm Heatmap, p Point, basin []Point) []Point {
  height := len(hm)
  width := len(hm[0])
  if !IsIn(basin, p) && hm[p.x][p.y] < 9 {
    basin = append(basin, p)
  }
  // Explore NSWE points recursively with DFS (not stask safe, but it works)
  left := Point{p.x, p.y-1}
  if left.y >= 0 && !IsIn(basin, left) && hm[left.x][left.y] < 9 {
    for _, newPoint := range Basin(hm, left, basin) {
      if !IsIn(basin, newPoint) {
        basin = append(basin, newPoint)
      }
    }
  }
  right := Point{p.x, p.y+1}
  if right.y < width && !IsIn(basin, right) && hm[right.x][right.y] < 9 {
    for _, newPoint := range Basin(hm, right, basin) {
      if !IsIn(basin, newPoint) {
        basin = append(basin, newPoint)
      }
    }
  }
  down := Point{p.x+1, p.y}
  if down.x < height && !IsIn(basin, down)  && hm[down.x][down.y] < 9 {
    for _, newPoint := range Basin(hm, down, basin) {
      if !IsIn(basin, newPoint) {
        basin = append(basin, newPoint)
      }
    }
  }
  up := Point{p.x-1, p.y}
  if up.x >= 0 && !IsIn(basin, up) && hm[up.x][up.y] < 9 {
    for _, newPoint := range Basin(hm, up, basin) {
      if !IsIn(basin, newPoint) {
        basin = append(basin, newPoint)
      }
    }
  }
  return basin
}

func Part2() {
  input, _ := utils.ReadLines("input.txt")
  hm := ParseInput(input)
  lp := hm.LowPoints()
  sizes := make([]int, 0)
  for _, p := range lp {
    basin := Basin(hm, p, make([]Point, 0))
    sizes = append(sizes, len(basin))
  }
  sort.Ints(sizes)
  prod := 1
  for _, s := range sizes[len(sizes)-3:] {
    prod *= s
  }
  fmt.Printf("Part 2: %d\n", prod)
}

func main() {
  Part1()
  Part2()
}
