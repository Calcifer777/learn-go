package main

import (
  "fmt"
  "strings"
  "strconv"
  "utils"
)

type Table [][]bool
type Point struct {
  x, y int
}
type Direction int64
const (
  Vertical   Direction = iota
  Horizontal 
)
type FoldOp struct {
  index int
  direction Direction
}

func ParseInput(lines []string) (Table, []FoldOp) {
  // Find table size
  points := make([]Point, 0)
  var maxX, maxY int
  foldIdx := 0
  for idx, line := range lines {
    if len(line) == 0 {
      foldIdx = idx
      break
    }
    coords := strings.Split(line, ",")
    x, _ := strconv.Atoi(coords[1])
    y, _ := strconv.Atoi(coords[0])
    if x > maxX { maxX = x }
    if y > maxY { maxY = y }
    points = append(points, Point{x, y})
  }
  // Create table
  table := make([][]bool, maxX+1)
  for r := 0; r <= maxX; r++ {
    table[r] = make([]bool, maxY+1)
  }
  for _, p := range points {
    table[p.x][p.y] = true
  }
  // Parse folds
  folds := make([]FoldOp, 0)
  for _, line := range lines[foldIdx+1:] {
    foldInfo := strings.Split(strings.Fields(line)[2], "=")
    foldIdx, _ := strconv.Atoi(foldInfo[1])
    var direction Direction
    switch foldInfo[0] {
    case "x": direction = Vertical
    case "y": direction = Horizontal
    default: panic("could not parse direction")
    }

    folds = append(folds, FoldOp{foldIdx, direction})
  }
  return table, folds
}

func (table Table) ToString() string {
  s := ""
  for r := 0; r < len(table); r++ {
    for c := 0; c < len(table[0]); c++ {
      if table[r][c] {
        s += "X"
      } else {
        s += " "
      }
    }
    s += "\n"
  }
  return s
}

func (table Table) Transpose() Table {
  transposed := make([][]bool, len(table[0]))
  for r := 0; r < len(table[0]); r++ {
    transposed[r] = make([]bool, len(table))
  }
  for r := 0; r < len(table); r++ {
    for c := 0; c < len(table[0]); c++ {
      transposed[c][r] = table[r][c]
    }
  }
  return transposed
}

// Always fold over the bottom sheet
func (table Table) HFold(fold FoldOp) Table {
  fold1 := table[:fold.index]    // top sheet
  fold2 := table[fold.index+1:]  // bottom sheet
  // fold1Len := fold.index
  fold2Len := len(fold2)
  colNum := len(table[0])
  folded := make(Table, 0)
  for r := 0; r < utils.Max(len(fold1), len(fold2)); r++ {
    if r < len(fold1) && r < len(fold2) {
      newLine := make([]bool, colNum)
      for c := 0; c < colNum; c++ {
        newLine[c] = fold2[fold2Len-r-1][c] || fold1[r][c]
      }
      folded = append(folded, newLine)
    } else if r < len(fold1) {
      folded = append(folded, fold1[r])
    } else if r < len(fold2) {
      folded = append(folded, fold2[fold2Len-r-1])
    }
  }
  return folded
}

// Leverage HFold; if vertical, fold the left sheet
func (table Table) Fold(foldOp FoldOp) Table {
  switch foldOp.direction {
  case Horizontal: {
    return table.HFold(foldOp)
  }
  case Vertical: {
    return table.Transpose().HFold(foldOp).Transpose()
  }
  default: panic("Unhandled fold direction")
  }
}

func (t *Table) Score() int {
  score := 0
  for _, r := range *t {
    for _, v := range r {
      if v {
        score += 1
      }
    }
  }
  return score
}

func Part1() {
  lines, _ := utils.ReadLines("input.txt")
  table, folds := ParseInput(lines)
  // fmt.Printf("%s\n\n", table.ToString())
  folded := table.Fold(folds[0])
  // fmt.Printf("%s\n", folded.ToString())
  fmt.Printf("Part 1 -> %d\n", folded.Score())
}

func Part2() {
  lines, _ := utils.ReadLines("input.txt")
  table, folds := ParseInput(lines)
  folded := table
  for _, fold := range folds {
    folded = folded.Fold(fold)
  }
  fmt.Printf("Part 2 ->\n%s\n", folded.ToString())
}

func main() {
  Part1()
  Part2()
}
